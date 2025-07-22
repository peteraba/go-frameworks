package service_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"github.com/peteraba/go-frameworks/shared/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListService_CRUD(t *testing.T) {
	r := repo.NewInMemoryListRepo()
	svc := service.NewListService(r)

	// Create
	lc := model.RandomListCreate()
	list, err := svc.Create(lc)
	require.NoError(t, err)
	assert.Equal(t, lc.Name, list.Name)
	assert.Equal(t, lc.Description, list.Description)
	assert.Equal(t, lc.ProjectID, list.ProjectID)
	assert.NotEmpty(t, list.ID)

	// GetByID
	got, err := svc.GetByID(list.ID)
	require.NoError(t, err)
	assert.Equal(t, list, got)

	// Update
	update := model.RandomListUpdate()
	updated, err := svc.Update(list.ID, update)
	require.NoError(t, err)
	assert.Equal(t, update.Name, updated.Name)
	assert.Equal(t, update.Description, updated.Description)

	// List
	lists, err := svc.List()
	require.NoError(t, err)
	assert.NotEmpty(t, lists)

	// Delete
	err = svc.Delete(list.ID)
	assert.NoError(t, err)
	_, err = svc.GetByID(list.ID)
	assert.Error(t, err)
}
