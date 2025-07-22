package service_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"github.com/peteraba/go-frameworks/shared/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTodoService_CRUD(t *testing.T) {
	repo := repo.NewInMemoryTodoRepo()
	svc := service.NewTodoService(repo)

	// Create
	tc := model.RandomTodoCreate()
	todo, err := svc.Create(tc)
	require.NoError(t, err)
	assert.Equal(t, tc.Title, todo.Title)
	assert.Equal(t, tc.Description, todo.Description)
	assert.Equal(t, tc.ListID, todo.ListID)
	assert.Equal(t, tc.Completed, todo.Completed)
	assert.NotEmpty(t, todo.ID)

	// GetByID
	got, err := svc.GetByID(todo.ID)
	require.NoError(t, err)
	assert.Equal(t, todo, got)

	// Update
	update := model.RandomTodoUpdate()
	updated, err := svc.Update(todo.ID, update)
	require.NoError(t, err)
	assert.Equal(t, update.Title, updated.Title)
	assert.Equal(t, update.Description, updated.Description)
	if update.Completed != nil {
		assert.Equal(t, *update.Completed, updated.Completed)
	}

	// List
	todos, err := svc.List()
	require.NoError(t, err)
	assert.NotEmpty(t, todos)

	// Delete
	err = svc.Delete(todo.ID)
	assert.NoError(t, err)
	_, err = svc.GetByID(todo.ID)
	assert.Error(t, err)
}
