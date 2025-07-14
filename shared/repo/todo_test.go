package repo_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryTodoRepo_Create(t *testing.T) {
	r := repo.NewInMemoryTodoRepo()

	t.Run("successful creation", func(t *testing.T) {
		// prepare
		todoCreateStub := model.RandomTodoCreate()
		todoCreateStub.ListID = "list-1"

		// execute
		todo, err := r.Create(todoCreateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, todoCreateStub.Title, todo.Title)
		assert.Equal(t, todoCreateStub.Description, todo.Description)
		assert.Equal(t, todoCreateStub.Completed, todo.Completed)
		assert.Equal(t, todoCreateStub.ListID, todo.ListID)
		assert.NotEmpty(t, todo.ID)
	})
}

func TestInMemoryTodoRepo_GetByID(t *testing.T) {
	r := repo.NewInMemoryTodoRepo()

	t.Run("existing todo", func(t *testing.T) {
		// prepare
		todoCreateStub := model.RandomTodoCreate()
		todoCreateStub.ListID = "list-1"
		todoStub, err := r.Create(todoCreateStub)
		require.NoError(t, err)
		retrievedStub, err := r.GetByID(todoStub.ID)

		// verify
		require.NoError(t, err)

		// execute
		assert.Equal(t, todoStub, retrievedStub)
	})

	t.Run("non-existing todo", func(t *testing.T) {
		_, err := r.GetByID("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrTodoNotFound)
	})
}

func TestInMemoryTodoRepo_Update(t *testing.T) {
	r := repo.NewInMemoryTodoRepo()

	t.Run("successful update", func(t *testing.T) {
		// prepare
		todoCreateStub := model.RandomTodoCreate()
		todoCreateStub.ListID = "list-1"
		todoStub, err := r.Create(todoCreateStub)
		require.NoError(t, err)
		todoUpdateStub := model.TodoUpdate{
			Title:       "Updated Title",
			Description: "Updated Description",
			Completed:   func() *bool { b := true; return &b }(),
		}

		// execute
		updated, err := r.Update(todoStub.ID, todoUpdateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Updated Title", updated.Title)
		assert.Equal(t, "Updated Description", updated.Description)
		assert.True(t, updated.Completed)
	})

	t.Run("partial update", func(t *testing.T) {
		// prepare
		todoCreateStub := model.RandomTodoCreate()
		todoCreateStub.ListID = "list-1"
		todoStub, err := r.Create(todoCreateStub)
		require.NoError(t, err)
		todoUpdateStub := model.TodoUpdate{
			Title: "Only Title Updated",
		}

		// execute
		updated, err := r.Update(todoStub.ID, todoUpdateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Only Title Updated", updated.Title)
		assert.Equal(t, todoStub.Description, updated.Description)
	})

	t.Run("non-existing todo", func(t *testing.T) {
		// prepare
		todoUpdateStub := model.TodoUpdate{Title: "Updated Title"}

		// execute
		_, err := r.Update("non-existing-id", todoUpdateStub)

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrTodoNotFound)
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		// prepare
		todoCreateStub := model.RandomTodoCreate()
		todoCreateStub.ListID = "list-1"
		todoStub, err := r.Create(todoCreateStub)
		require.NoError(t, err)
		todoUpdateStub := model.TodoUpdate{}

		// execute
		updated, err := r.Update(todoStub.ID, todoUpdateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, todoStub.Title, updated.Title)
		assert.Equal(t, todoStub.Description, updated.Description)
		assert.Equal(t, todoStub.Completed, updated.Completed)
	})
}

func TestInMemoryTodoRepo_Delete(t *testing.T) {
	r := repo.NewInMemoryTodoRepo()

	t.Run("successful deletion", func(t *testing.T) {
		// prepare
		todoCreateStub := model.RandomTodoCreate()
		todoCreateStub.ListID = "list-1"
		todoStub, err := r.Create(todoCreateStub)
		require.NoError(t, err)

		// execute
		err = r.Delete(todoStub.ID)

		// verify
		assert.NoError(t, err)
	})

	t.Run("non-existing todo", func(t *testing.T) {
		// execute
		err := r.Delete("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrTodoNotFound)
	})
}

func TestInMemoryTodoRepo_List(t *testing.T) {
	r := repo.NewInMemoryTodoRepo()

	t.Run("empty repo", func(t *testing.T) {
		// execute
		todos, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.NotNil(t, todos)
		assert.Equal(t, 0, len(todos))
	})

	t.Run("with todos", func(t *testing.T) {
		// prepare
		todoCreateStub1 := model.RandomTodoCreate()
		todoCreateStub1.ListID = "list-1"
		todoCreateStub2 := model.RandomTodoCreate()
		todoCreateStub2.ListID = "list-1"
		todoStub1, _ := r.Create(todoCreateStub1)
		todoStub2, _ := r.Create(todoCreateStub2)

		// execute
		todos, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 2, len(todos))
		assert.True(t, r.Has(todoStub1.ID))
		assert.True(t, r.Has(todoStub2.ID))
	})

	t.Run("with 1000+ todos", func(t *testing.T) {
		// prepare
		for range 1005 {
			todoCreateStub := model.RandomTodoCreate()
			_, err := r.Create(todoCreateStub)
			require.NoError(t, err)
		}

		// execute
		lists, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 1000, len(lists))
	})
}
