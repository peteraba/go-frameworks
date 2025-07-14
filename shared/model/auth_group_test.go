package model_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/stretchr/testify/assert"
)

func TestRandomAuthGroup(t *testing.T) {
	// execute
	got1 := model.RandomAuthGroup()
	got2 := model.RandomAuthGroup()

	// verify
	assert.NotEmpty(t, got1.ID)
	assert.NotEmpty(t, got1.Name)
	assert.NotEqual(t, got1.ID, got2.ID)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomAuthGroupCreate(t *testing.T) {
	// execute
	got1 := model.RandomAuthGroupCreate()
	got2 := model.RandomAuthGroupCreate()

	// verify
	assert.NotEmpty(t, got1.Name)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomAuthGroupUpdate(t *testing.T) {
	// execute
	got1 := model.RandomAuthGroupUpdate()
	got2 := model.RandomAuthGroupUpdate()

	// verify
	assert.NotEmpty(t, got1.Name)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}
