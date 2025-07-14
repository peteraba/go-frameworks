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
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"

		// execute
		todo, err := r.Create(tc)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, tc.Title, todo.Title)
		assert.Equal(t, tc.Description, todo.Description)
		assert.Equal(t, tc.Completed, todo.Completed)
		assert.Equal(t, tc.ListID, todo.ListID)
		assert.NotEmpty(t, todo.ID)
	})
}

func TestInMemoryTodoRepo_GetByID(t *testing.T) {
	r := repo.NewInMemoryTodoRepo()

	t.Run("existing todo", func(t *testing.T) {
		// prepare
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := r.Create(tc)
		require.NoError(t, err)
		retrieved, err := r.GetByID(todo.ID)

		// verify
		require.NoError(t, err)

		// execute
		assert.Equal(t, todo, retrieved)
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
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := r.Create(tc)
		require.NoError(t, err)
		update := model.TodoUpdate{
			Title:       "Updated Title",
			Description: "Updated Description",
			Completed:   func() *bool { b := true; return &b }(),
		}

		// execute
		updated, err := r.Update(todo.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Updated Title", updated.Title)
		assert.Equal(t, "Updated Description", updated.Description)
		assert.True(t, updated.Completed)
	})

	t.Run("partial update", func(t *testing.T) {
		// prepare
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := r.Create(tc)
		require.NoError(t, err)
		update := model.TodoUpdate{
			Title: "Only Title Updated",
		}

		// execute
		updated, err := r.Update(todo.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Only Title Updated", updated.Title)
		assert.Equal(t, todo.Description, updated.Description)
	})

	t.Run("non-existing todo", func(t *testing.T) {
		// prepare
		update := model.TodoUpdate{Title: "Updated Title"}

		// execute
		_, err := r.Update("non-existing-id", update)

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrTodoNotFound)
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		// prepare
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := r.Create(tc)
		require.NoError(t, err)
		update := model.TodoUpdate{}

		// execute
		updated, err := r.Update(todo.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, todo.Title, updated.Title)
		assert.Equal(t, todo.Description, updated.Description)
		assert.Equal(t, todo.Completed, updated.Completed)
	})
}

func TestInMemoryTodoRepo_Delete(t *testing.T) {
	r := repo.NewInMemoryTodoRepo()

	t.Run("successful deletion", func(t *testing.T) {
		// prepare
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := r.Create(tc)
		require.NoError(t, err)

		// execute
		err = r.Delete(todo.ID)

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
		tc1 := model.RandomTodoCreate()
		tc1.ListID = "list-1"
		tc2 := model.RandomTodoCreate()
		tc2.ListID = "list-1"
		todo1, _ := r.Create(tc1)
		todo2, _ := r.Create(tc2)

		// execute
		todos, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 2, len(todos))
		assert.True(t, r.Has(todo1.ID))
		assert.True(t, r.Has(todo2.ID))
	})
}
