package service_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"github.com/peteraba/go-frameworks/shared/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectService_CRUD(t *testing.T) {
	repo := repo.NewInMemoryProjectRepo()
	svc := service.NewProjectService(repo)

	// Create
	pc := model.RandomProjectCreate()
	project, err := svc.Create(pc)
	require.NoError(t, err)
	assert.Equal(t, pc.Name, project.Name)
	assert.Equal(t, pc.Description, project.Description)
	assert.NotEmpty(t, project.ID)

	// GetByID
	got, err := svc.GetByID(project.ID)
	require.NoError(t, err)
	assert.Equal(t, project, got)

	// Update
	update := model.RandomProjectUpdate()
	updated, err := svc.Update(project.ID, update)
	require.NoError(t, err)
	assert.Equal(t, update.Name, updated.Name)
	assert.Equal(t, update.Description, updated.Description)

	// List
	projects, err := svc.List()
	require.NoError(t, err)
	assert.NotEmpty(t, projects)

	// Delete
	err = svc.Delete(project.ID)
	assert.NoError(t, err)
	_, err = svc.GetByID(project.ID)
	assert.Error(t, err)
}
