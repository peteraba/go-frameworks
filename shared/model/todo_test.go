package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomTodo(t *testing.T) {
	got1 := RandomTodo()
	got2 := RandomTodo()

	assert.NotEmpty(t, got1)
	assert.NotEmpty(t, got2)
	assert.NotEqual(t, got1, got2)
}

func TestRandomTodoCreate(t *testing.T) {
	got1 := RandomTodoCreate()
	got2 := RandomTodoCreate()

	assert.NotEmpty(t, got1)
	assert.NotEmpty(t, got2)
	assert.NotEqual(t, got1, got2)
}

func TestRandomTodoUpdate(t *testing.T) {
	got1 := RandomTodoUpdate()
	got2 := RandomTodoUpdate()

	assert.NotEmpty(t, got1)
	assert.NotEmpty(t, got2)
	assert.NotEqual(t, got1, got2)
}
