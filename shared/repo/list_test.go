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
		// prepare
		listCreateStub := model.RandomListCreate()

		// execute
		created, err := r.Create(listCreateStub)

		// verify
		assert.NoError(t, err)

		assert.Equal(t, listCreateStub.Name, created.Name)
		assert.Equal(t, listCreateStub.Description, created.Description)
		assert.True(t, r.Has(created.ID))
	})
}

func TestInMemoryListRepo_GetByID(t *testing.T) {
	r := repo.NewInMemoryListRepo()

	t.Run("existing list", func(t *testing.T) {
		// prepare
		listCreateStub := model.RandomListCreate()

		listStub, err := r.Create(model.ListCreate{
			ProjectID:   listCreateStub.ProjectID,
			Name:        listCreateStub.Name,
			Description: listCreateStub.Description,
		})
		require.NoError(t, err)

		// execute
		retrieved, err := r.GetByID(listStub.ID)

		// verify
		assert.NoError(t, err)

		assert.Equal(t, listStub, retrieved)
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
		// prepare
		listToCreateStub := model.RandomListCreate()
		createdListStub, err := r.Create(listToCreateStub)
		require.NoError(t, err)

		listUpdateStub := model.RandomListUpdate()

		// execute
		updated, err := r.Update(createdListStub.ID, listUpdateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, listUpdateStub.Name, updated.Name)
		assert.Equal(t, listUpdateStub.Description, updated.Description)
		assert.Equal(t, createdListStub.ID, updated.ID)
		assert.Equal(t, createdListStub.ProjectID, updated.ProjectID)

		// Verify the list was actually updated in the repo
		retrieved, err := r.GetByID(updated.ID)
		require.NoError(t, err)
		assert.Equal(t, listUpdateStub.Name, retrieved.Name)
		assert.Equal(t, listUpdateStub.Description, retrieved.Description)
	})

	t.Run("non-existing list", func(t *testing.T) {
		// prepare
		listUpdateStub := model.ListUpdate{
			Name: "Updated Name",
		}

		// execute
		_, err := r.Update("non-existing-id", listUpdateStub)

		// verify
		assert.Error(t, err)
		assert.Equal(t, "list not found", err.Error())
	})
}

func TestInMemoryListRepo_Delete(t *testing.T) {
	r := repo.NewInMemoryListRepo()

	t.Run("successful deletion", func(t *testing.T) {
		// prepare
		listCreateStub := model.RandomListCreate()

		listStub, err := r.Create(listCreateStub)
		require.NoError(t, err)

		// execute
		err = r.Delete(listStub.ID)

		// verify
		assert.NoError(t, err)

		assert.False(t, r.Has(listStub.ID))
	})

	t.Run("non-existing list", func(t *testing.T) {
		// execute
		err := r.Delete("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.Equal(t, "list not found", err.Error())
	})
}

func TestInMemoryListRepo_List(t *testing.T) {
	r := repo.NewInMemoryListRepo()

	t.Run("empty repo", func(t *testing.T) {
		// execute
		lists, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.NotNil(t, lists)
		assert.Empty(t, lists)
	})

	t.Run("with lists", func(t *testing.T) {
		// prepare
		listCreateStub1 := model.RandomListCreate()
		listCreateStub2 := model.RandomListCreate()

		listStub1, err := r.Create(listCreateStub1)
		require.NoError(t, err)
		listStub2, err := r.Create(listCreateStub2)
		require.NoError(t, err)

		// execute
		lists, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 2, len(lists))
		assert.True(t, r.Has(listStub1.ID))
		assert.True(t, r.Has(listStub2.ID))
	})

	t.Run("with 100+ lists", func(t *testing.T) {
		// prepare
		for range 105 {
			listCreateStub := model.RandomListCreate()
			_, err := r.Create(listCreateStub)
			require.NoError(t, err)
		}

		// execute
		lists, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 100, len(lists))
	})
}

func TestInMemoryListRepo_Concurrency(t *testing.T) {
	r := repo.NewInMemoryListRepo()

	t.Run("concurrent create and read", func(t *testing.T) {
		// prepare
		done := make(chan bool)

		// goroutine 1: Create lists
		go func() {
			for range 100 {
				list := model.RandomListCreate()
				_, err := r.Create(list)
				require.NoError(t, err)
			}
			done <- true
		}()

		// goroutine 2: Read lists
		go func() {
			for range 100 {
				_, err := r.List()
				require.NoError(t, err)
			}
			done <- true
		}()

		<-done
		<-done

		// execute
		lists, err := r.List()

		// verify that no panic occurred and data is consistent
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(lists), 100) // May be less due to duplicate ID errors
	})
}
