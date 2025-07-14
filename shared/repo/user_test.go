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
		ucStub := model.RandomUserCreate()
		passwordStub := []byte{}

		// execute
		user, err := r.Create(ucStub, passwordStub)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, ucStub.Name, user.Name)
		assert.Equal(t, ucStub.Email, user.Email)
		assert.NotEmpty(t, user.ID)
	})
}

func TestInMemoryUserRepo_GetByID(t *testing.T) {
	r := repo.NewInMemoryUserRepo()

	t.Run("existing user", func(t *testing.T) {
		// prepare
		ucStub := model.RandomUserCreate()
		passwordStub := []byte{}

		user, err := r.Create(ucStub, passwordStub)
		require.NoError(t, err)

		// execute
		retrieved, err := r.GetByID(user.ID)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, user, retrieved)
	})

	t.Run("non-existing user", func(t *testing.T) {
		// execute
		_, err := r.GetByID("non-existing-id")

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrUserNotFound)
	})
}

func TestInMemoryUserRepo_Update(t *testing.T) {
	r := repo.NewInMemoryUserRepo()

	t.Run("successful update", func(t *testing.T) {
		// prepare
		ucStub := model.RandomUserCreate()
		passwordStub := []byte{}

		user, err := r.Create(ucStub, passwordStub)
		require.NoError(t, err)
		update := model.UserUpdate{
			Name:  "Updated Name",
			Email: "updated@example.com",
		}

		// execute
		updated, err := r.Update(user.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updated.Name)
		assert.Equal(t, "updated@example.com", updated.Email)
		assert.Equal(t, user.ID, updated.ID)
	})

	t.Run("partial update", func(t *testing.T) {
		// prepare
		ucStub := model.RandomUserCreate()
		passwordStub := []byte{}

		user, err := r.Create(ucStub, passwordStub)
		require.NoError(t, err)
		update := model.UserUpdate{
			Name: "Only Name Updated",
		}

		// execute
		updated, err := r.Update(user.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, "Only Name Updated", updated.Name)
		assert.Equal(t, user.Email, updated.Email)
	})

	t.Run("non-existing user", func(t *testing.T) {
		// prepare
		uuStub := model.RandomUserUpdate()

		// execute
		_, err := r.Update("non-existing-id", uuStub)

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrUserNotFound)
	})

	t.Run("empty update fields are ignored", func(t *testing.T) {
		// prepare
		ucStub := model.RandomUserCreate()
		passwordStub := []byte{}

		user, err := r.Create(ucStub, passwordStub)
		require.NoError(t, err)
		update := model.UserUpdate{}

		// execute
		updated, err := r.Update(user.ID, update)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, user.Name, updated.Name)
		assert.Equal(t, user.Email, updated.Email)
	})
}

func TestInMemoryUserRepo_Delete(t *testing.T) {
	r := repo.NewInMemoryUserRepo()

	t.Run("successful deletion", func(t *testing.T) {
		// prepare
		ucStub := model.RandomUserCreate()
		passwordStub := []byte{}

		user, err := r.Create(ucStub, passwordStub)
		require.NoError(t, err)

		// execute
		err = r.Delete(user.ID)

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
		uc1Stub := model.RandomUserCreate()
		uc2Stub := model.RandomUserCreate()
		passwordStub := []byte{}

		user1, _ := r.Create(uc1Stub, passwordStub)
		user2, _ := r.Create(uc2Stub, passwordStub)

		// execute
		users, err := r.List()

		// verify
		assert.NoError(t, err)
		assert.Equal(t, 2, len(users))
		assert.True(t, r.Has(user1.ID))
		assert.True(t, r.Has(user2.ID))
	})
}
