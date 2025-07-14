package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-playground/validator/v10"
)

type Todo struct {
	ID          string `json:"id" validate:"required,max=26" fake:"{ulid}"`
	ListID      string `json:"listId" validate:"required,max=26" fake:"{ulid}"`
	Title       string `json:"title" validate:"required,max=64" fake:"{sentence:3}"`
	Description string `json:"description,omitempty" validate:"max=255" fake:"{paragraph:1}"`
	Completed   bool   `json:"completed"`
}

type TodoCreate struct {
	ListID      string `json:"listId" validate:"required,max=26"`
	Title       string `json:"title" validate:"required,max=64"`
	Description string `json:"description,omitempty" validate:"max=255"`
	Completed   bool   `json:"completed"`
}

type TodoUpdate struct {
	Title       string `json:"title,omitempty" validate:"max=64"`
	Description string `json:"description,omitempty" validate:"max=255"`
	Completed   *bool  `json:"completed,omitempty"`
}

var validate = validator.New()

// Todo validation methods
func (t *Todo) Validate() error {
	return validate.Struct(t)
}

func (tc *TodoCreate) Validate() error {
	return validate.Struct(tc)
}

func (tu *TodoUpdate) Validate() error {
	return validate.Struct(tu)
}

// Random generation methods for Todo
func RandomTodo() Todo {
	var t Todo

	err := gofakeit.Struct(&t)
	if err != nil {
		panic(err)
	}

	return t
}

func RandomTodoCreate() TodoCreate {
	t := RandomTodo()

	return TodoCreate{
		ListID:      t.ListID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
	}
}

func RandomTodoUpdate() TodoUpdate {
	t := RandomTodo()
	completed := t.Completed

	return TodoUpdate{
		Title:       t.Title,
		Description: t.Description,
		Completed:   &completed,
	}
}
