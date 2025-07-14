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
		agc := model.RandomAuthGroupCreate()
		ag, err := repo.Create(agc)
		require.NoError(t, err)
		assert.Equal(t, agc.Name, ag.Name)
		assert.ElementsMatch(t, agc.Users, ag.Users)
		assert.NotEmpty(t, ag.ID)
	})
}

func TestInMemoryAuthGroupRepo_GetByID(t *testing.T) {
	repo := NewInMemoryAuthGroupRepo()

	t.Run("existing group", func(t *testing.T) {
		agc := model.RandomAuthGroupCreate()
		ag, err := repo.Create(agc)
		require.NoError(t, err)
		retrieved, err := repo.GetByID(ag.ID)
		require.NoError(t, err)
		assert.Equal(t, ag, retrieved)
	})

	t.Run("non-existing group", func(t *testing.T) {
		_, err := repo.GetByID("non-existing-id")
		assert.Error(t, err)
		assert.Equal(t, "auth group not found", err.Error())
	})
}

func TestInMemoryAuthGroupRepo_Update(t *testing.T) {
	repo := NewInMemoryAuthGroupRepo()

	t.Run("successful update", func(t *testing.T) {
		agc := model.RandomAuthGroupCreate()
		ag, err := repo.Create(agc)
		require.NoError(t, err)
		update := model.AuthGroupUpdate{
			Name:  "Updated Name",
			Users: []string{"user1", "user2"},
		}
		updated, err := repo.Update(ag.ID, update)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updated.Name)
		assert.ElementsMatch(t, []string{"user1", "user2"}, updated.Users)
		assert.Equal(t, ag.ID, updated.ID)
	})

	t.Run("partial update", func(t *testing.T) {
		agc := model.RandomAuthGroupCreate()
		ag, err := repo.Create(agc)
		require.NoError(t, err)
		update := model.AuthGroupUpdate{
			Name: "Only Name Updated",
		}
		updated, err := repo.Update(ag.ID, update)
		assert.NoError(t, err)
		assert.Equal(t, "Only Name Updated", updated.Name)
		assert.ElementsMatch(t, ag.Users, updated.Users)
	})

	t.Run("non-existing group", func(t *testing.T) {
		update := model.AuthGroupUpdate{Name: "Updated Name"}
		_, err := repo.Update("non-existing-id", update)
		assert.Error(t, err)
		assert.Equal(t, "auth group not found", err.Error())
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		agc := model.RandomAuthGroupCreate()
		ag, err := repo.Create(agc)
		require.NoError(t, err)
		update := model.AuthGroupUpdate{}
		updated, err := repo.Update(ag.ID, update)
		assert.NoError(t, err)
		assert.Equal(t, ag.Name, updated.Name)
		assert.ElementsMatch(t, ag.Users, updated.Users)
	})
}

func TestInMemoryAuthGroupRepo_Delete(t *testing.T) {
	repo := NewInMemoryAuthGroupRepo()

	t.Run("successful deletion", func(t *testing.T) {
		agc := model.RandomAuthGroupCreate()
		ag, err := repo.Create(agc)
		require.NoError(t, err)
		err = repo.Delete(ag.ID)
		assert.NoError(t, err)
	})

	t.Run("non-existing group", func(t *testing.T) {
		err := repo.Delete("non-existing-id")
		assert.Error(t, err)
		assert.Equal(t, "auth group not found", err.Error())
	})
}

func TestInMemoryAuthGroupRepo_List(t *testing.T) {
	repo := NewInMemoryAuthGroupRepo()

	t.Run("empty repo", func(t *testing.T) {
		groups, err := repo.List()
		assert.NoError(t, err)
		assert.NotNil(t, groups)
		assert.Equal(t, 0, len(groups))
	})

	t.Run("with groups", func(t *testing.T) {
		agc1 := model.RandomAuthGroupCreate()
		agc2 := model.RandomAuthGroupCreate()
		ag1, _ := repo.Create(agc1)
		ag2, _ := repo.Create(agc2)
		groups, err := repo.List()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(groups))
		ids := map[string]bool{ag1.ID: false, ag2.ID: false}
		for _, g := range groups {
			ids[g.ID] = true
		}
		assert.True(t, ids[ag1.ID])
		assert.True(t, ids[ag2.ID])
	})
}
