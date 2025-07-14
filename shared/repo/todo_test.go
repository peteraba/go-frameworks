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
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := repo.Create(tc)
		require.NoError(t, err)
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
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := repo.Create(tc)
		require.NoError(t, err)
		retrieved, err := repo.GetByID(todo.ID)
		require.NoError(t, err)
		assert.Equal(t, todo, retrieved)
	})

	t.Run("non-existing todo", func(t *testing.T) {
		_, err := repo.GetByID("non-existing-id")
		assert.Error(t, err)
		assert.Equal(t, "todo not found", err.Error())
	})
}

func TestInMemoryTodoRepo_Update(t *testing.T) {
	repo := NewInMemoryTodoRepo()

	t.Run("successful update", func(t *testing.T) {
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := repo.Create(tc)
		require.NoError(t, err)
		update := model.TodoUpdate{
			Title:       "Updated Title",
			Description: "Updated Description",
			Completed:   func() *bool { b := true; return &b }(),
		}
		updated, err := repo.Update(todo.ID, update)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Title", updated.Title)
		assert.Equal(t, "Updated Description", updated.Description)
		assert.True(t, updated.Completed)
	})

	t.Run("partial update", func(t *testing.T) {
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := repo.Create(tc)
		require.NoError(t, err)
		update := model.TodoUpdate{
			Title: "Only Title Updated",
		}
		updated, err := repo.Update(todo.ID, update)
		assert.NoError(t, err)
		assert.Equal(t, "Only Title Updated", updated.Title)
		assert.Equal(t, todo.Description, updated.Description)
	})

	t.Run("non-existing todo", func(t *testing.T) {
		update := model.TodoUpdate{Title: "Updated Title"}
		_, err := repo.Update("non-existing-id", update)
		assert.Error(t, err)
		assert.Equal(t, "todo not found", err.Error())
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := repo.Create(tc)
		require.NoError(t, err)
		update := model.TodoUpdate{}
		updated, err := repo.Update(todo.ID, update)
		assert.NoError(t, err)
		assert.Equal(t, todo.Title, updated.Title)
		assert.Equal(t, todo.Description, updated.Description)
		assert.Equal(t, todo.Completed, updated.Completed)
	})
}

func TestInMemoryTodoRepo_Delete(t *testing.T) {
	repo := NewInMemoryTodoRepo()

	t.Run("successful deletion", func(t *testing.T) {
		tc := model.RandomTodoCreate()
		tc.ListID = "list-1"
		todo, err := repo.Create(tc)
		require.NoError(t, err)
		err = repo.Delete(todo.ID)
		assert.NoError(t, err)
	})

	t.Run("non-existing todo", func(t *testing.T) {
		err := repo.Delete("non-existing-id")
		assert.Error(t, err)
		assert.Equal(t, "todo not found", err.Error())
	})
}

func TestInMemoryTodoRepo_List(t *testing.T) {
	repo := NewInMemoryTodoRepo()

	t.Run("empty repo", func(t *testing.T) {
		todos, err := repo.List()
		assert.NoError(t, err)
		assert.NotNil(t, todos)
		assert.Equal(t, 0, len(todos))
	})

	t.Run("with todos", func(t *testing.T) {
		tc1 := model.RandomTodoCreate()
		tc1.ListID = "list-1"
		tc2 := model.RandomTodoCreate()
		tc2.ListID = "list-1"
		todo1, _ := repo.Create(tc1)
		todo2, _ := repo.Create(tc2)
		todos, err := repo.List()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(todos))
		ids := map[string]bool{todo1.ID: false, todo2.ID: false}
		for _, t := range todos {
			ids[t.ID] = true
		}
		assert.True(t, ids[todo1.ID])
		assert.True(t, ids[todo2.ID])
	})
}
