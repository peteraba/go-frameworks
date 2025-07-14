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
	assert.NotEmpty(t, got1.Name)
	assert.NotEqual(t, got1.ID, got2.ID)
	assert.NotEqual(t, got1.Name, got2.Name)
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
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}
