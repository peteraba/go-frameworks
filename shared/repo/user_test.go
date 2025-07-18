package repo_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryUserRepo_Create(t *testing.T) {
	r := repo.NewInMemoryUserRepo()

	t.Run("successful creation", func(t *testing.T) {
		// prepare
		userCreateStub := model.RandomUserCreate()
		passwordHashStub := []byte{}
		passwordSaltStub := []byte{}

		// execute
		user, err := r.Create(userCreateStub, passwordHashStub, passwordSaltStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, userCreateStub.Name, user.Name)
		assert.Equal(t, userCreateStub.Email, user.Email)
		assert.NotEmpty(t, user.ID)
	})
}

func TestInMemoryUserRepo_GetByID(t *testing.T) {
	r := repo.NewInMemoryUserRepo()

	t.Run("existing user", func(t *testing.T) {
		// prepare
		userCreateStub := model.RandomUserCreate()
		passwordHashStub := []byte{}
		passwordSaltStub := []byte{}

		userStub, err := r.Create(userCreateStub, passwordHashStub, passwordSaltStub)
		require.NoError(t, err)

		// execute
		retrieved, err := r.GetByID(userStub.ID)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, userStub, retrieved)
	})

	t.Run("non-existing user", func(t *testing.T) {
		// execute
		_, err := r.GetByID("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrUserNotFound)
	})
}

func TestInMemoryUserRepo_GetByEmail(t *testing.T) {
	r := repo.NewInMemoryUserRepo()

	t.Run("existing user by email", func(t *testing.T) {
		// prepare
		userCreateStub := model.RandomUserCreate()
		passwordHashStub := []byte{}
		passwordSaltStub := []byte{}

		userStub, err := r.Create(userCreateStub, passwordHashStub, passwordSaltStub)
		require.NoError(t, err)

		// execute
		retrieved, err := r.GetByEmail(userStub.Email)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, userStub, retrieved)
	})

	t.Run("non-existing email", func(t *testing.T) {
		// execute
		_, err := r.GetByEmail("nonexistent@example.com")

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrUserNotFound)
	})

	t.Run("case sensitive email matching", func(t *testing.T) {
		// prepare
		userCreateStub := model.RandomUserCreate()
		userCreateStub.Email = "TestUser@Example.com"
		passwordHashStub := []byte{}
		passwordSaltStub := []byte{}

		userStub, err := r.Create(userCreateStub, passwordHashStub, passwordSaltStub)
		require.NoError(t, err)

		// execute - should not find with different case
		_, err = r.GetByEmail("testuser@example.com")

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrUserNotFound)

		// execute - should find with exact case
		retrieved, err := r.GetByEmail("TestUser@Example.com")

		// verify
		assert.NoError(t, err)
		assert.Equal(t, userStub, retrieved)
	})

	t.Run("multiple users with different emails", func(t *testing.T) {
		// prepare
		userCreateStub1 := model.RandomUserCreate()
		userCreateStub2 := model.RandomUserCreate()
		passwordHashStub := []byte{}
		passwordSaltStub := []byte{}

		userStub1, err := r.Create(userCreateStub1, passwordHashStub, passwordSaltStub)
		require.NoError(t, err)
		userStub2, err := r.Create(userCreateStub2, passwordHashStub, passwordSaltStub)
		require.NoError(t, err)

		// execute
		retrieved1, err := r.GetByEmail(userCreateStub1.Email)
		assert.NoError(t, err)
		retrieved2, err := r.GetByEmail(userCreateStub2.Email)
		assert.NoError(t, err)

		// verify
		assert.Equal(t, userStub1, retrieved1)
		assert.Equal(t, userStub2, retrieved2)
		assert.NotEqual(t, retrieved1.ID, retrieved2.ID)
	})
}

func TestInMemoryUserRepo_Update(t *testing.T) {
	r := repo.NewInMemoryUserRepo()

	t.Run("successful update", func(t *testing.T) {
		// prepare
		userCreateStub := model.RandomUserCreate()
		passwordHashStub := []byte{}
		passwordSaltStub := []byte{}

		userStub, err := r.Create(userCreateStub, passwordHashStub, passwordSaltStub)
		require.NoError(t, err)
		userUpdateStub := model.UserUpdate{
			Name:  "Updated Name",
			Email: "updated@example.com",
		}

		// execute
		updated, err := r.Update(userStub.ID, userUpdateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updated.Name)
		assert.Equal(t, "updated@example.com", updated.Email)
		assert.Equal(t, userStub.ID, updated.ID)
	})

	t.Run("partial update", func(t *testing.T) {
		// prepare
		userCreateStub := model.RandomUserCreate()
		passwordHashStub := []byte{}
		passwordSaltStub := []byte{}

		userStub, err := r.Create(userCreateStub, passwordHashStub, passwordSaltStub)
		require.NoError(t, err)
		userUpdateStub := model.UserUpdate{
			Name: "Only Name Updated",
		}

		// execute
		updated, err := r.Update(userStub.ID, userUpdateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Only Name Updated", updated.Name)
		assert.Equal(t, userStub.Email, updated.Email)
	})

	t.Run("non-existing user", func(t *testing.T) {
		// prepare
		userUpdateStub := model.RandomUserUpdate()

		// execute
		_, err := r.Update("non-existing-id", userUpdateStub)

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrUserNotFound)
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		// prepare
		userCreateStub := model.RandomUserCreate()
		passwordHashStub := []byte{}
		passwordSaltStub := []byte{}

		userStub, err := r.Create(userCreateStub, passwordHashStub, passwordSaltStub)
		require.NoError(t, err)
		userUpdateStub := model.UserUpdate{}

		// execute
		updated, err := r.Update(userStub.ID, userUpdateStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, userStub.Name, updated.Name)
		assert.Equal(t, userStub.Email, updated.Email)
	})
}

func TestInMemoryUserRepo_Delete(t *testing.T) {
	r := repo.NewInMemoryUserRepo()

	t.Run("successful deletion", func(t *testing.T) {
		// prepare
		userCreateStub := model.RandomUserCreate()
		passwordHashStub := []byte{}
		passwordSaltStub := []byte{}

		userStub, err := r.Create(userCreateStub, passwordHashStub, passwordSaltStub)
		require.NoError(t, err)

		// execute
		err = r.Delete(userStub.ID)

		// verify
		assert.NoError(t, err)
	})

	t.Run("non-existing user", func(t *testing.T) {
		// execute
		err := r.Delete("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrUserNotFound)
	})
}

func TestInMemoryUserRepo_List(t *testing.T) {
	r := repo.NewInMemoryUserRepo()

	t.Run("empty repo", func(t *testing.T) {
		// execute
		users, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, 0, len(users))
	})

	t.Run("with users", func(t *testing.T) {
		// prepare
		userCreateStub1 := model.RandomUserCreate()
		userCreateStub2 := model.RandomUserCreate()
		passwordHashStub := []byte{}
		passwordSaltStub := []byte{}

		userStub1, _ := r.Create(userCreateStub1, passwordHashStub, passwordSaltStub)
		userStub2, _ := r.Create(userCreateStub2, passwordHashStub, passwordSaltStub)

		// execute
		users, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 2, len(users))
		assert.True(t, r.Has(userStub1.ID))
		assert.True(t, r.Has(userStub2.ID))
	})

	t.Run("with 100+ lists", func(t *testing.T) {
		var passwordHashStub, passwordSaltStub []byte

		// prepare
		for range 105 {
			userCreateStub := model.RandomUserCreate()
			_, err := r.Create(userCreateStub, passwordHashStub, passwordSaltStub)
			require.NoError(t, err)
		}

		// execute
		lists, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 100, len(lists))
	})
}
