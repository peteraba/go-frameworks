package repo

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryAuthGroupRepo_Create(t *testing.T) {
	repo := NewInMemoryAuthGroupRepo()

	t.Run("successful creation", func(t *testing.T) {
		// prepare
		agc := model.RandomAuthGroupCreate()

		// execute
		ag, err := repo.Create(agc)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, agc.Name, ag.Name)
		assert.NotEmpty(t, ag.ID)
	})
}

func TestInMemoryAuthGroupRepo_GetByID(t *testing.T) {
	repo := NewInMemoryAuthGroupRepo()

	t.Run("existing group", func(t *testing.T) {
		// prepare
		agc := model.RandomAuthGroupCreate()

		ag, err := repo.Create(agc)
		require.NoError(t, err)

		// execute
		retrieved, err := repo.GetByID(ag.ID)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, ag, retrieved)
	})

	t.Run("non-existing group", func(t *testing.T) {
		// execute
		_, err := repo.GetByID("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.Equal(t, "auth group not found", err.Error())
	})
}

func TestInMemoryAuthGroupRepo_Update(t *testing.T) {
	repo := NewInMemoryAuthGroupRepo()

	t.Run("successful update", func(t *testing.T) {
		// prepare
		agc := model.RandomAuthGroupCreate()
		ag, err := repo.Create(agc)
		require.NoError(t, err)
		update := model.AuthGroupUpdate{
			Name: "Updated Name",
		}

		// execute
		updated, err := repo.Update(ag.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updated.Name)
		assert.Equal(t, ag.ID, updated.ID)
	})

	t.Run("partial update", func(t *testing.T) {
		// prepare
		agc := model.RandomAuthGroupCreate()
		ag, err := repo.Create(agc)
		require.NoError(t, err)
		update := model.AuthGroupUpdate{
			Name: "Only Name Updated",
		}

		// execute
		updated, err := repo.Update(ag.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Only Name Updated", updated.Name)
	})

	t.Run("non-existing group", func(t *testing.T) {
		// prepare
		update := model.AuthGroupUpdate{Name: "Updated Name"}

		// execute
		_, err := repo.Update("non-existing-id", update)

		// verify
		assert.Error(t, err)
		assert.Equal(t, "auth group not found", err.Error())
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		// prepare
		agc := model.RandomAuthGroupCreate()
		ag, err := repo.Create(agc)
		require.NoError(t, err)
		update := model.AuthGroupUpdate{}

		// execute
		updated, err := repo.Update(ag.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, ag.Name, updated.Name)
	})
}

func TestInMemoryAuthGroupRepo_Delete(t *testing.T) {
	repo := NewInMemoryAuthGroupRepo()

	t.Run("successful deletion", func(t *testing.T) {
		// prepare
		agc := model.RandomAuthGroupCreate()

		ag, err := repo.Create(agc)
		require.NoError(t, err)

		// execute
		err = repo.Delete(ag.ID)

		// verify
		assert.NoError(t, err)

		ag2, err := repo.GetByID(ag.ID)
		assert.Error(t, err)
		assert.Empty(t, ag2)
	})

	t.Run("non-existing group", func(t *testing.T) {
		// execute
		err := repo.Delete("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.Equal(t, "auth group not found", err.Error())
	})
}

func TestInMemoryAuthGroupRepo_List(t *testing.T) {
	repo := NewInMemoryAuthGroupRepo()

	t.Run("empty repo", func(t *testing.T) {
		// execute
		groups, err := repo.List()

		// verify
		assert.NoError(t, err)
		assert.NotNil(t, groups)
		assert.Equal(t, 0, len(groups))
	})

	t.Run("with groups", func(t *testing.T) {
		// prepare
		agc1 := model.RandomAuthGroupCreate()
		agc2 := model.RandomAuthGroupCreate()
		ag1, _ := repo.Create(agc1)
		ag2, _ := repo.Create(agc2)

		// execute
		groups, err := repo.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 2, len(groups))
		assert.True(t, repo.Has(ag1.ID))
		assert.True(t, repo.Has(ag2.ID))
	})
}
