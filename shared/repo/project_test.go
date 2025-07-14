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
		projectCreateStub := model.RandomProjectCreate()

		// execute
		project, err := r.Create(projectCreateStub)

		// verify
		require.NoError(t, err)
		assert.Equal(t, projectCreateStub.Name, project.Name)
		assert.Equal(t, projectCreateStub.Description, project.Description)
		assert.NotEmpty(t, project.ID)
	})
}

func TestInMemoryProjectRepo_GetByID(t *testing.T) {
	r := repo.NewInMemoryProjectRepo()

	t.Run("existing project", func(t *testing.T) {
		// prepare
		projectCreateStub := model.RandomProjectCreate()
		projectStub, err := r.Create(projectCreateStub)
		require.NoError(t, err)

		// execute
		retrieved, err := r.GetByID(projectStub.ID)

		// verify
		require.NoError(t, err)
		assert.Equal(t, projectStub, retrieved)
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
		projectCreateStub := model.RandomProjectCreate()
		projectStub, err := r.Create(projectCreateStub)
		require.NoError(t, err)
		projectUpdateStub := model.ProjectUpdate{
			Name:        "Updated Name",
			Description: "Updated Description",
		}

		// execute
		updated, err := r.Update(projectStub.ID, projectUpdateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updated.Name)
		assert.Equal(t, "Updated Description", updated.Description)
		assert.Equal(t, projectStub.ID, updated.ID)
	})

	t.Run("partial update", func(t *testing.T) {
		// prepare
		projectCreateStub := model.RandomProjectCreate()
		projectStub, err := r.Create(projectCreateStub)
		require.NoError(t, err)
		projectUpdateStub := model.ProjectUpdate{
			Name: "Only Name Updated",
		}

		// execute
		updated, err := r.Update(projectStub.ID, projectUpdateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Only Name Updated", updated.Name)
		assert.Equal(t, projectStub.Description, updated.Description)
	})

	t.Run("non-existing project", func(t *testing.T) {
		// prepare
		projectUpdateStub := model.ProjectUpdate{Name: "Updated Name"}

		// execute
		_, err := r.Update("non-existing-id", projectUpdateStub)

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrProjectNotFound)
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		// prepare
		projectCreateStub := model.RandomProjectCreate()
		projectStub, err := r.Create(projectCreateStub)
		require.NoError(t, err)
		projectUpdateStub := model.ProjectUpdate{}

		// execute
		updated, err := r.Update(projectStub.ID, projectUpdateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, projectStub.Name, updated.Name)
		assert.Equal(t, projectStub.Description, updated.Description)
	})
}

func TestInMemoryProjectRepo_Delete(t *testing.T) {
	r := repo.NewInMemoryProjectRepo()

	t.Run("successful deletion", func(t *testing.T) {
		// prepare
		projectCreateStub := model.RandomProjectCreate()
		projectStub, err := r.Create(projectCreateStub)
		require.NoError(t, err)

		// execute
		err = r.Delete(projectStub.ID)

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
		projectCreateStub1 := model.RandomProjectCreate()
		projectCreateStub2 := model.RandomProjectCreate()
		projectStub1, _ := r.Create(projectCreateStub1)
		projectStub2, _ := r.Create(projectCreateStub2)

		// execute
		projects, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 2, len(projects))
		assert.True(t, r.Has(projectStub1.ID))
		assert.True(t, r.Has(projectStub2.ID))
	})

	t.Run("with 100+ lists", func(t *testing.T) {
		// prepare
		for range 105 {
			projectCreateStub := model.RandomProjectCreate()
			_, err := r.Create(projectCreateStub)
			require.NoError(t, err)
		}

		// execute
		lists, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 100, len(lists))
	})
}
