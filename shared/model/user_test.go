package model_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/stretchr/testify/assert"
)

func TestRandomUser(t *testing.T) {
	// execute
	got1 := model.RandomUser()
	got2 := model.RandomUser()

	// verify
	assert.NotEmpty(t, got1.ID)
	assert.NotEqual(t, got1.ID, got2.ID)
	assert.NotEmpty(t, got1.Name)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEmpty(t, got1.Groups)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomUserCreate(t *testing.T) {
	// execute
	got1 := model.RandomUserCreate()
	got2 := model.RandomUserCreate()

	// verify
	assert.NotEmpty(t, got1.Name)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEmpty(t, got1.Email)
	assert.NotEqual(t, got1.Email, got2.Email)
	assert.NotEmpty(t, got1.Groups)
	assert.NotEmpty(t, got1.Password)
	assert.NotEqual(t, got1.Password, got2.Password)
	assert.Equal(t, got1.Password, got1.Password2)
	assert.Equal(t, got2.Password, got2.Password2)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomUserUpdate(t *testing.T) {
	// execute
	got1 := model.RandomUserUpdate()
	got2 := model.RandomUserUpdate()

	// verify
	assert.NotEmpty(t, got1.Name)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEmpty(t, got1.Email)
	assert.NotEqual(t, got1.Email, got2.Email)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomUserLogin(t *testing.T) {
	// execute
	got1 := model.RandomUserLogin()
	got2 := model.RandomUserLogin()

	// verify
	assert.NotEmpty(t, got1.Name)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEmpty(t, got1.Password)
	assert.NotEqual(t, got1.Password, got2.Password)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomUserPasswordUpdate(t *testing.T) {
	// execute
	got1 := model.RandomUserPasswordUpdate()
	got2 := model.RandomUserPasswordUpdate()

	// verify
	assert.NotEmpty(t, got1.Password)
	assert.NotEqual(t, got1.Password, got2.Password)
	assert.Equal(t, got1.Password, got1.Password2)
	assert.Equal(t, got2.Password, got2.Password2)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}
