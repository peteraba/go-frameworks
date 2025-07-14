package repo_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryProjectRepo_Create(t *testing.T) {
	r := repo.NewInMemoryProjectRepo()

	t.Run("successful creation", func(t *testing.T) {
		// prepare
		pc := model.RandomProjectCreate()

		// execute
		project, err := r.Create(pc)

		// verify
		require.NoError(t, err)
		assert.Equal(t, pc.Name, project.Name)
		assert.Equal(t, pc.Description, project.Description)
		assert.NotEmpty(t, project.ID)
	})
}

func TestInMemoryProjectRepo_GetByID(t *testing.T) {
	r := repo.NewInMemoryProjectRepo()

	t.Run("existing project", func(t *testing.T) {
		// prepare
		pc := model.RandomProjectCreate()
		project, err := r.Create(pc)
		require.NoError(t, err)

		// execute
		retrieved, err := r.GetByID(project.ID)

		// verify
		require.NoError(t, err)
		assert.Equal(t, project, retrieved)
	})

	t.Run("non-existing project", func(t *testing.T) {
		// execute
		_, err := r.GetByID("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrProjectNotFound)
	})
}

func TestInMemoryProjectRepo_Update(t *testing.T) {
	r := repo.NewInMemoryProjectRepo()

	t.Run("successful update", func(t *testing.T) {
		// prepare
		pc := model.RandomProjectCreate()
		project, err := r.Create(pc)
		require.NoError(t, err)
		update := model.ProjectUpdate{
			Name:        "Updated Name",
			Description: "Updated Description",
		}

		// execute
		updated, err := r.Update(project.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updated.Name)
		assert.Equal(t, "Updated Description", updated.Description)
		assert.Equal(t, project.ID, updated.ID)
	})

	t.Run("partial update", func(t *testing.T) {
		// prepare
		pc := model.RandomProjectCreate()
		project, err := r.Create(pc)
		require.NoError(t, err)
		update := model.ProjectUpdate{
			Name: "Only Name Updated",
		}

		// execute
		updated, err := r.Update(project.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Only Name Updated", updated.Name)
		assert.Equal(t, project.Description, updated.Description)
	})

	t.Run("non-existing project", func(t *testing.T) {
		// prepare
		update := model.ProjectUpdate{Name: "Updated Name"}

		// execute
		_, err := r.Update("non-existing-id", update)

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrProjectNotFound)
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		// prepare
		pc := model.RandomProjectCreate()
		project, err := r.Create(pc)
		require.NoError(t, err)
		update := model.ProjectUpdate{}

		// execute
		updated, err := r.Update(project.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, project.Name, updated.Name)
		assert.Equal(t, project.Description, updated.Description)
	})
}

func TestInMemoryProjectRepo_Delete(t *testing.T) {
	r := repo.NewInMemoryProjectRepo()

	t.Run("successful deletion", func(t *testing.T) {
		// prepare
		pc := model.RandomProjectCreate()
		project, err := r.Create(pc)
		require.NoError(t, err)

		// execute
		err = r.Delete(project.ID)

		// verify
		assert.NoError(t, err)
	})

	t.Run("non-existing project", func(t *testing.T) {
		// execute
		err := r.Delete("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrProjectNotFound)
	})
}

func TestInMemoryProjectRepo_List(t *testing.T) {
	r := repo.NewInMemoryProjectRepo()

	t.Run("empty repo", func(t *testing.T) {
		// execute
		projects, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.NotNil(t, projects)
		assert.Empty(t, projects)
	})

	t.Run("with projects", func(t *testing.T) {
		// prepare
		pc1 := model.RandomProjectCreate()
		pc2 := model.RandomProjectCreate()
		project1, _ := r.Create(pc1)
		project2, _ := r.Create(pc2)

		// execute
		projects, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 2, len(projects))
		assert.True(t, r.Has(project1.ID))
		assert.True(t, r.Has(project2.ID))
	})
}
