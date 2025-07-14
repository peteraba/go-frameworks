package repo

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryUserRepo_Create(t *testing.T) {
	repo := NewInMemoryUserRepo()

	t.Run("successful creation", func(t *testing.T) {
		// prepare
		uc := model.RandomUserCreate()

		// execute
		user, err := repo.Create(uc)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, uc.Name, user.Name)
		assert.Equal(t, uc.Email, user.Email)
		assert.NotEmpty(t, user.ID)
	})
}

func TestInMemoryUserRepo_GetByID(t *testing.T) {
	repo := NewInMemoryUserRepo()

	t.Run("existing user", func(t *testing.T) {
		// prepare
		uc := model.RandomUserCreate()
		user, err := repo.Create(uc)
		require.NoError(t, err)

		// execute
		retrieved, err := repo.GetByID(user.ID)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, user, retrieved)
	})

	t.Run("non-existing user", func(t *testing.T) {
		// execute
		_, err := repo.GetByID("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestInMemoryUserRepo_Update(t *testing.T) {
	repo := NewInMemoryUserRepo()

	t.Run("successful update", func(t *testing.T) {
		// prepare
		uc := model.RandomUserCreate()
		user, err := repo.Create(uc)
		require.NoError(t, err)
		update := model.UserUpdate{
			Name:  "Updated Name",
			Email: "updated@example.com",
		}

		// execute
		updated, err := repo.Update(user.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updated.Name)
		assert.Equal(t, "updated@example.com", updated.Email)
		assert.Equal(t, user.ID, updated.ID)
	})

	t.Run("partial update", func(t *testing.T) {
		// prepare
		uc := model.RandomUserCreate()
		user, err := repo.Create(uc)
		require.NoError(t, err)
		update := model.UserUpdate{
			Name: "Only Name Updated",
		}

		// execute
		updated, err := repo.Update(user.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Only Name Updated", updated.Name)
		assert.Equal(t, user.Email, updated.Email)
	})

	t.Run("non-existing user", func(t *testing.T) {
		// prepare
		update := model.UserUpdate{Name: "Updated Name"}

		// execute
		_, err := repo.Update("non-existing-id", update)

		// verify
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		// prepare
		uc := model.RandomUserCreate()
		user, err := repo.Create(uc)
		require.NoError(t, err)
		update := model.UserUpdate{}

		// execute
		updated, err := repo.Update(user.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, user.Name, updated.Name)
		assert.Equal(t, user.Email, updated.Email)
	})
}

func TestInMemoryUserRepo_Delete(t *testing.T) {
	repo := NewInMemoryUserRepo()

	t.Run("successful deletion", func(t *testing.T) {
		// prepare
		uc := model.RandomUserCreate()
		user, err := repo.Create(uc)
		require.NoError(t, err)

		// execute
		err = repo.Delete(user.ID)

		// verify
		assert.NoError(t, err)
	})

	t.Run("non-existing user", func(t *testing.T) {
		// execute
		err := repo.Delete("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestInMemoryUserRepo_List(t *testing.T) {
	repo := NewInMemoryUserRepo()

	t.Run("empty repo", func(t *testing.T) {
		// execute
		users, err := repo.List()

		// verify
		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, 0, len(users))
	})

	t.Run("with users", func(t *testing.T) {
		// prepare
		uc1 := model.RandomUserCreate()
		uc2 := model.RandomUserCreate()
		user1, _ := repo.Create(uc1)
		user2, _ := repo.Create(uc2)

		// execute
		users, err := repo.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 2, len(users))
		assert.True(t, repo.Has(user1.ID))
		assert.True(t, repo.Has(user2.ID))
	})
}
