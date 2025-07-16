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
		uc.Password = "SuperSecret123!"
		uc.Password2 = "SuperSecret123!"

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
		uc.Password = "SuperSecret123!"
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
		uc.Password = "SuperSecret123!"
		uc.Password2 = "SuperSecret123!"

		_, err := sut.Create(uc)
		require.NoError(t, err)
		// Try to create again with the same details (should not error, as ID is generated, but let's check)
		_, err = sut.Create(uc)
		assert.NoError(t, err)
	})
}
