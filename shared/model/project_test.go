package model_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/stretchr/testify/assert"
)

func TestRandomProject(t *testing.T) {
	// execute
	got1 := model.RandomProject()
	got2 := model.RandomProject()

	// verify
	assert.NotEmpty(t, got1.ID)
	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.ID, got2.ID)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomProjectCreate(t *testing.T) {
	// execute
	got1 := model.RandomProjectCreate()
	got2 := model.RandomProjectCreate()

	// verify
	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomProjectUpdate(t *testing.T) {
	// execute
	got1 := model.RandomProjectUpdate()
	got2 := model.RandomProjectUpdate()

	// verify
	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}
