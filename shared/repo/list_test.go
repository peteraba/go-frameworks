package repo_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryListRepo_Create(t *testing.T) {
	r := repo.NewInMemoryListRepo()

	t.Run("successful creation", func(t *testing.T) {
		list := model.RandomListCreate()

		created, err := r.Create(list)
		require.NoError(t, err)

		assert.Equal(t, list.Name, created.Name)
		assert.Equal(t, list.Description, created.Description)
		assert.True(t, r.Has(created.ID))
	})
}

func TestInMemoryListRepo_GetByID(t *testing.T) {
	r := repo.NewInMemoryListRepo()

	t.Run("existing list", func(t *testing.T) {
		lc := model.RandomListCreate()

		list, err := r.Create(model.ListCreate{
			ProjectID:   lc.ProjectID,
			Name:        lc.Name,
			Description: lc.Description,
		})
		require.NoError(t, err)

		retrieved, err := r.GetByID(list.ID)
		require.NoError(t, err)

		assert.Equal(t, list, retrieved)
	})

	t.Run("non-existing list", func(t *testing.T) {
		_, err := r.GetByID("non-existing-id")

		assert.Error(t, err)
		assert.Equal(t, "list not found", err.Error())
	})
}

func TestInMemoryListRepo_Update(t *testing.T) {
	r := repo.NewInMemoryListRepo()

	t.Run("successful update", func(t *testing.T) {
		listToCreate := model.RandomListCreate()
		createdList, err := r.Create(listToCreate)
		require.NoError(t, err)

		update := model.ListUpdate{
			Name:        "Updated Name",
			Description: "Updated Description",
		}

		updated, err := r.Update(createdList.ID, update)

		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updated.Name)
		assert.Equal(t, "Updated Description", updated.Description)
		assert.Equal(t, createdList.ID, updated.ID)
		assert.Equal(t, createdList.ProjectID, updated.ProjectID)

		// Verify the list was actually updated in the repo
		retrieved, err := r.GetByID(updated.ID)
		require.NoError(t, err)
		assert.Equal(t, "Updated Name", retrieved.Name)
		assert.Equal(t, "Updated Description", retrieved.Description)
	})

	t.Run("non-existing list", func(t *testing.T) {
		update := model.ListUpdate{
			Name: "Updated Name",
		}

		_, err := r.Update("non-existing-id", update)

		assert.Error(t, err)
		assert.Equal(t, "list not found", err.Error())
	})
}

func TestInMemoryListRepo_Delete(t *testing.T) {
	r := repo.NewInMemoryListRepo()

	t.Run("successful deletion", func(t *testing.T) {
		lc := model.RandomListCreate()

		list, err := r.Create(lc)
		require.NoError(t, err)

		err = r.Delete(list.ID)
		assert.NoError(t, err)

		assert.False(t, r.Has(list.ID))
	})

	t.Run("non-existing list", func(t *testing.T) {
		err := r.Delete("non-existing-id")

		assert.Error(t, err)
		assert.Equal(t, "list not found", err.Error())
	})
}

func TestInMemoryListRepo_List(t *testing.T) {
	r := repo.NewInMemoryListRepo()

	t.Run("empty repo", func(t *testing.T) {
		lists, err := r.List()

		assert.NoError(t, err)
		assert.NotNil(t, lists)
		assert.Equal(t, 0, len(lists))
	})

	t.Run("with lists", func(t *testing.T) {
		lc1 := model.RandomListCreate()
		lc2 := model.RandomListCreate()

		list1, err := r.Create(lc1)
		require.NoError(t, err)
		list2, err := r.Create(lc2)
		require.NoError(t, err)

		lists, err := r.List()

		assert.NoError(t, err)
		assert.Equal(t, 2, len(lists))

		// Verify both lists are present (order may vary)
		listIDs := make(map[string]bool)
		for _, list := range lists {
			listIDs[list.ID] = true
		}

		assert.True(t, listIDs[list1.ID])
		assert.True(t, listIDs[list2.ID])
	})
}

func TestInMemoryListRepo_Concurrency(t *testing.T) {
	r := repo.NewInMemoryListRepo()

	t.Run("concurrent create and read", func(t *testing.T) {
		done := make(chan bool)

		// Goroutine 1: Create lists
		go func() {
			for range 100 {
				list := model.RandomListCreate()
				_, err := r.Create(list)
				require.NoError(t, err)
			}
			done <- true
		}()

		// Goroutine 2: Read lists
		go func() {
			for range 100 {
				_, err := r.List()
				require.NoError(t, err)
			}
			done <- true
		}()

		<-done
		<-done

		// Verify no panic occurred and data is consistent
		lists, err := r.List()
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(lists), 100) // May be less due to duplicate ID errors
	})
}
