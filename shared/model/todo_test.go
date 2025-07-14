package model_test

import (
	"testing"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/stretchr/testify/assert"
)

func TestRandomTodo(t *testing.T) {
	// execute
	got1 := model.RandomTodo()
	got2 := model.RandomTodo()

	// verify
	assert.NotEmpty(t, got1.ID)
	assert.NotEmpty(t, got1.Title)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Title, got2.Title)
	assert.NotEqual(t, got1.Description, got2.Description)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomTodoCreate(t *testing.T) {
	// execute
	got1 := model.RandomTodoCreate()
	got2 := model.RandomTodoCreate()

	// verify
	assert.NotEmpty(t, got1.Title)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Title, got2.Title)
	assert.NotEqual(t, got1.Description, got2.Description)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}

func TestRandomTodoUpdate(t *testing.T) {
	// execute
	got1 := model.RandomTodoUpdate()
	got2 := model.RandomTodoUpdate()

	// verify
	assert.NotEmpty(t, got1.Title)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Title, got2.Title)
	assert.NotEqual(t, got1.Description, got2.Description)
	assert.NoError(t, got1.Validate())
	assert.NoError(t, got2.Validate())
}
