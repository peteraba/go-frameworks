package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/peteraba/go-frameworks/shared/model"
)

func TestRandomList(t *testing.T) {
	// execute
	got1 := model.RandomList()
	got2 := model.RandomList()

	// verify
	assert.NotEmpty(t, got1.ID)
	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.ID, got2.ID)
	assert.NotEqual(t, got1.ProjectID, got2.ProjectID)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomListCreate(t *testing.T) {
	// execute
	got1 := model.RandomListCreate()
	got2 := model.RandomListCreate()

	// verify
	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.ProjectID, got2.ProjectID)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomListUpdate(t *testing.T) {
	// execute
	got1 := model.RandomListUpdate()
	got2 := model.RandomListUpdate()

	// verify
	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}
