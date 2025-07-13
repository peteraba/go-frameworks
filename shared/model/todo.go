package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-playground/validator/v10"
)

type Todo struct {
	ID          string `json:"id" validate:"required,max=36" fake:"{ulid}"`
	Title       string `json:"title" validate:"required,max=64,alphanumunicode" fake:"{sentence:3}"`
	Description string `json:"description,omitempty" validate:"max=255,alphanumunicode" fake:"{paragraph:1}"`
	Completed   bool   `json:"completed"`
}

type TodoCreate struct {
	Title       string `json:"title" validate:"required,max=64,alphanumunicode"`
	Description string `json:"description,omitempty" validate:"max=255,alphanumunicode"`
}

type TodoUpdate struct {
	Title       string `json:"title,omitempty" validate:"max=64,alphanumunicode"`
	Description string `json:"description,omitempty" validate:"max=255,alphanumunicode"`
	Completed   bool   `json:"completed,omitempty"`
}

var validate = validator.New()

func (t *Todo) Validate() error {
	return validate.Struct(t)
}

func (tc *TodoCreate) Validate() error {
	return validate.Struct(tc)
}

func (tu *TodoUpdate) Validate() error {
	return validate.Struct(tu)
}

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
		Title:       t.Title,
		Description: t.Description,
	}
}

func RandomTodoUpdate() TodoUpdate {
	t := RandomTodo()

	return TodoUpdate{
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
	}
}
