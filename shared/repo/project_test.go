package repo

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryProjectRepo_Create(t *testing.T) {
	repo := NewInMemoryProjectRepo()

	t.Run("successful creation", func(t *testing.T) {
		pc := model.RandomProjectCreate()
		project, err := repo.Create(pc)
		require.NoError(t, err)
		assert.Equal(t, pc.Name, project.Name)
		assert.Equal(t, pc.Description, project.Description)
		assert.NotEmpty(t, project.ID)
	})
}

func TestInMemoryProjectRepo_GetByID(t *testing.T) {
	repo := NewInMemoryProjectRepo()

	t.Run("existing project", func(t *testing.T) {
		pc := model.RandomProjectCreate()
		project, err := repo.Create(pc)
		require.NoError(t, err)
		retrieved, err := repo.GetByID(project.ID)
		require.NoError(t, err)
		assert.Equal(t, project, retrieved)
	})

	t.Run("non-existing project", func(t *testing.T) {
		_, err := repo.GetByID("non-existing-id")
		assert.Error(t, err)
		assert.Equal(t, "project not found", err.Error())
	})
}

func TestInMemoryProjectRepo_Update(t *testing.T) {
	repo := NewInMemoryProjectRepo()

	t.Run("successful update", func(t *testing.T) {
		pc := model.RandomProjectCreate()
		project, err := repo.Create(pc)
		require.NoError(t, err)
		update := model.ProjectUpdate{
			Name:        "Updated Name",
			Description: "Updated Description",
		}
		updated, err := repo.Update(project.ID, update)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updated.Name)
		assert.Equal(t, "Updated Description", updated.Description)
		assert.Equal(t, project.ID, updated.ID)
	})

	t.Run("partial update", func(t *testing.T) {
		pc := model.RandomProjectCreate()
		project, err := repo.Create(pc)
		require.NoError(t, err)
		update := model.ProjectUpdate{
			Name: "Only Name Updated",
		}
		updated, err := repo.Update(project.ID, update)
		assert.NoError(t, err)
		assert.Equal(t, "Only Name Updated", updated.Name)
		assert.Equal(t, project.Description, updated.Description)
	})

	t.Run("non-existing project", func(t *testing.T) {
		update := model.ProjectUpdate{Name: "Updated Name"}
		_, err := repo.Update("non-existing-id", update)
		assert.Error(t, err)
		assert.Equal(t, "project not found", err.Error())
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		pc := model.RandomProjectCreate()
		project, err := repo.Create(pc)
		require.NoError(t, err)
		update := model.ProjectUpdate{}
		updated, err := repo.Update(project.ID, update)
		assert.NoError(t, err)
		assert.Equal(t, project.Name, updated.Name)
		assert.Equal(t, project.Description, updated.Description)
	})
}

func TestInMemoryProjectRepo_Delete(t *testing.T) {
	repo := NewInMemoryProjectRepo()

	t.Run("successful deletion", func(t *testing.T) {
		pc := model.RandomProjectCreate()
		project, err := repo.Create(pc)
		require.NoError(t, err)
		err = repo.Delete(project.ID)
		assert.NoError(t, err)
	})

	t.Run("non-existing project", func(t *testing.T) {
		err := repo.Delete("non-existing-id")
		assert.Error(t, err)
		assert.Equal(t, "project not found", err.Error())
	})
}

func TestInMemoryProjectRepo_List(t *testing.T) {
	repo := NewInMemoryProjectRepo()

	t.Run("empty repo", func(t *testing.T) {
		projects, err := repo.List()
		assert.NoError(t, err)
		assert.NotNil(t, projects)
		assert.Equal(t, 0, len(projects))
	})

	t.Run("with projects", func(t *testing.T) {
		pc1 := model.RandomProjectCreate()
		pc2 := model.RandomProjectCreate()
		project1, _ := repo.Create(pc1)
		project2, _ := repo.Create(pc2)
		projects, err := repo.List()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(projects))
		ids := map[string]bool{project1.ID: false, project2.ID: false}
		for _, p := range projects {
			ids[p.ID] = true
		}
		assert.True(t, ids[project1.ID])
		assert.True(t, ids[project2.ID])
	})
}
