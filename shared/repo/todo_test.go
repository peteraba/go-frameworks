package repo

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryTodoRepo_Create(t *testing.T) {
	repo := NewInMemoryTodoRepo()

	t.Run("successful creation", func(t *testing.T) {
		// prepare
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"

		// execute
		todo, err := repo.Create(tc)

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
	repo := NewInMemoryTodoRepo()

	t.Run("existing todo", func(t *testing.T) {
		// prepare
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := repo.Create(tc)
		require.NoError(t, err)
		retrieved, err := repo.GetByID(todo.ID)

		// verify
		require.NoError(t, err)

		// execute
		assert.Equal(t, todo, retrieved)
	})

	t.Run("non-existing todo", func(t *testing.T) {
		_, err := repo.GetByID("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.Equal(t, "todo not found", err.Error())
	})
}

func TestInMemoryTodoRepo_Update(t *testing.T) {
	repo := NewInMemoryTodoRepo()

	t.Run("successful update", func(t *testing.T) {
		// prepare
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := repo.Create(tc)
		require.NoError(t, err)
		update := model.TodoUpdate{
			Title:       "Updated Title",
			Description: "Updated Description",
			Completed:   func() *bool { b := true; return &b }(),
		}

		// execute
		updated, err := repo.Update(todo.ID, update)

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
		todo, err := repo.Create(tc)
		require.NoError(t, err)
		update := model.TodoUpdate{
			Title: "Only Title Updated",
		}

		// execute
		updated, err := repo.Update(todo.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Only Title Updated", updated.Title)
		assert.Equal(t, todo.Description, updated.Description)
	})

	t.Run("non-existing todo", func(t *testing.T) {
		// prepare
		update := model.TodoUpdate{Title: "Updated Title"}

		// execute
		_, err := repo.Update("non-existing-id", update)

		// verify
		assert.Error(t, err)
		assert.Equal(t, "todo not found", err.Error())
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		// prepare
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := repo.Create(tc)
		require.NoError(t, err)
		update := model.TodoUpdate{}

		// execute
		updated, err := repo.Update(todo.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, todo.Title, updated.Title)
		assert.Equal(t, todo.Description, updated.Description)
		assert.Equal(t, todo.Completed, updated.Completed)
	})
}

func TestInMemoryTodoRepo_Delete(t *testing.T) {
	repo := NewInMemoryTodoRepo()

	t.Run("successful deletion", func(t *testing.T) {
		// prepare
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := repo.Create(tc)
		require.NoError(t, err)

		// execute
		err = repo.Delete(todo.ID)

		// verify
		assert.NoError(t, err)
	})

	t.Run("non-existing todo", func(t *testing.T) {
		// execute
		err := repo.Delete("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.Equal(t, "todo not found", err.Error())
	})
}

func TestInMemoryTodoRepo_List(t *testing.T) {
	repo := NewInMemoryTodoRepo()

	t.Run("empty repo", func(t *testing.T) {
		// execute
		todos, err := repo.List()

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
		todo1, _ := repo.Create(tc1)
		todo2, _ := repo.Create(tc2)

		// execute
		todos, err := repo.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 2, len(todos))
		assert.True(t, repo.Has(todo1.ID))
		assert.True(t, repo.Has(todo2.ID))
	})
}
