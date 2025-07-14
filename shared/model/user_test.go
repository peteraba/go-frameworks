package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation(t *testing.T) {
	u := RandomUser()
	assert.NoError(t, u.Validate())

	u.ID = ""
	assert.Error(t, u.Validate())

	u = RandomUser()
	u.Email = "not-an-email"
	assert.Error(t, u.Validate())
}

func TestUserCreateValidation(t *testing.T) {
	uc := RandomUserCreate()
	assert.NoError(t, uc.Validate())

	uc.Name = ""
	assert.Error(t, uc.Validate())

	uc = RandomUserCreate()
	uc.Email = "not-an-email"
	assert.Error(t, uc.Validate())
}

func TestUserUpdateValidation(t *testing.T) {
	uu := RandomUserUpdate()
	assert.NoError(t, uu.Validate())

	uu.Email = "not-an-email"
	assert.Error(t, uu.Validate())
}

func TestRandomUserGeneration(t *testing.T) {
	u := RandomUser()
	assert.NotEmpty(t, u.ID)
	assert.NotEmpty(t, u.Name)
	assert.NotEmpty(t, u.Email)
}
