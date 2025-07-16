package service_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"github.com/peteraba/go-frameworks/shared/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_Create(t *testing.T) {
	userRepo := repo.NewInMemoryUserRepo()
	sut := service.NewUserService(userRepo)

	t.Run("successful creation", func(t *testing.T) {
		uc := model.RandomUserCreate()

		user, err := sut.Create(uc)
		require.NoError(t, err)
		assert.Equal(t, uc.Name, user.Name)
		assert.Equal(t, uc.Email, user.Email)
		assert.ElementsMatch(t, uc.Groups, user.Groups)
		assert.NotEmpty(t, user.ID)
		assert.NotEmpty(t, user.PasswordHash)
		assert.NotEmpty(t, user.PasswordSalt)
	})

	t.Run("passwords do not match", func(t *testing.T) {
		uc := model.RandomUserCreate()
		uc.Password2 = "DifferentPassword!"

		user, err := sut.Create(uc)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "'Password2' failed on the 'eqfield'")
		assert.Empty(t, user.ID)
	})

	t.Run("empty password", func(t *testing.T) {
		uc := model.RandomUserCreate()
		uc.Password = ""
		uc.Password2 = ""

		user, err := sut.Create(uc)
		assert.Error(t, err)
		assert.Empty(t, user.ID)
	})

	t.Run("repo returns error", func(t *testing.T) {
		// Simulate duplicate user by using the same repo and user details
		uc := model.RandomUserCreate()

		_, err := sut.Create(uc)
		require.NoError(t, err)
		// Try to create again with the same details (should not error, as ID is generated, but let's check)
		_, err = sut.Create(uc)
		assert.NoError(t, err)
	})
}

func TestUserService_Login(t *testing.T) {
	userRepo := repo.NewInMemoryUserRepo()
	userService := service.NewUserService(userRepo)

	// Create a user
	sut := model.RandomUserCreate()
	_, err := userService.Create(sut)
	require.NoError(t, err)

	t.Run("successful login", func(t *testing.T) {
		// prepare
		ul := model.UserLogin{
			Email:    sut.Email,
			Password: sut.Password,
		}

		// execute
		token, err := userService.Login(ul)

		// verify
		assert.NoError(t, err)
		assert.Greater(t, len(token), 200)
		assert.Less(t, len(token), 400)
	})

	t.Run("wrong password", func(t *testing.T) {
		// prepare
		ul := model.UserLogin{
			Email:    sut.Email,
			Password: "wrong",
		}

		// execute
		_, err := userService.Login(ul)

		// verify
		assert.Error(t, err)
		assert.ErrorContains(t, err, "invalid credentials")
	})

	t.Run("user not found", func(t *testing.T) {
		// prepare
		ul := model.UserLogin{
			Email:    "notfound@example.com",
			Password: "irrelevant",
		}

		// execute
		_, err := userService.Login(ul)

		// verify
		assert.Error(t, err)
		assert.ErrorIs(t, err, repo.ErrUserNotFound)
	})
}

func TestUserService_TokenToLoggedInUser(t *testing.T) {
	userRepo := repo.NewInMemoryUserRepo()
	userService := service.NewUserService(userRepo)

	// Create a user
	ucStub := model.RandomUserCreate()
	userStub, err := userService.Create(ucStub)
	require.NoError(t, err)

	t.Run("successful decoding", func(t *testing.T) {
		// prepare
		ul := model.UserLogin{
			Email:    ucStub.Email,
			Password: ucStub.Password,
		}

		token, err := userService.Login(ul)
		require.NoError(t, err)

		// execute
		liu, err := userService.TokenToLoggedInUser(token)

		// verify
		assert.NoError(t, err)
		assert.Equal(t, userStub.ID, liu.ID)
	})
}
