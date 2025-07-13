package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	InitFaker()
}

func TestRandomProject(t *testing.T) {

	got1 := RandomProject()
	got2 := RandomProject()

	assert.NotEmpty(t, got1.ID)
	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.ID, got2.ID)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
}

func TestRandomProjectCreate(t *testing.T) {
	got1 := RandomProjectCreate()
	got2 := RandomProjectCreate()

	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
}

func TestRandomProjectUpdate(t *testing.T) {
	got1 := RandomProjectUpdate()
	got2 := RandomProjectUpdate()

	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
}

func TestRandomList(t *testing.T) {
	got1 := RandomList()
	got2 := RandomList()

	assert.NotEmpty(t, got1.ID)
	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.ID, got2.ID)
	assert.NotEqual(t, got1.ProjectID, got2.ProjectID)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
}

func TestRandomListCreate(t *testing.T) {
	got1 := RandomListCreate()
	got2 := RandomListCreate()

	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.ProjectID, got2.ProjectID)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
}

func TestRandomListUpdate(t *testing.T) {
	got1 := RandomListUpdate()
	got2 := RandomListUpdate()

	assert.NotEmpty(t, got1.Name)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Name, got2.Name)
	assert.NotEqual(t, got1.Description, got2.Description)
}

func TestRandomTodo(t *testing.T) {
	got1 := RandomTodo()
	got2 := RandomTodo()

	assert.NotEmpty(t, got1.ID)
	assert.NotEmpty(t, got1.Title)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Title, got2.Title)
	assert.NotEqual(t, got1.Description, got2.Description)
}

func TestRandomTodoCreate(t *testing.T) {
	got1 := RandomTodoCreate()
	got2 := RandomTodoCreate()

	assert.NotEmpty(t, got1.Title)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Title, got2.Title)
	assert.NotEqual(t, got1.Description, got2.Description)
}

func TestRandomTodoUpdate(t *testing.T) {
	got1 := RandomTodoUpdate()
	got2 := RandomTodoUpdate()

	assert.NotEmpty(t, got1.Title)
	assert.NotEmpty(t, got1.Description)
	assert.NotEqual(t, got1.Title, got2.Title)
	assert.NotEqual(t, got1.Description, got2.Description)
}
