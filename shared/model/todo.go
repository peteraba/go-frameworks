package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-playground/validator/v10"
)

type Project struct {
	ID          string `json:"id" validate:"required,max=26" fake:"{ulid}"`
	Name        string `json:"name" validate:"required,max=64" fake:"{sentence:2}"`
	Description string `json:"description,omitempty" validate:"max=64" fake:"{sentence:4}"`
}

type ProjectCreate struct {
	Name        string `json:"name" validate:"required,max=64"`
	Description string `json:"description,omitempty" validate:"max=255"`
}

type ProjectUpdate struct {
	Name        string `json:"name,omitempty" validate:"max=64"`
	Description string `json:"description,omitempty" validate:"max=255"`
}

type List struct {
	ID          string `json:"id" validate:"required,max=26" fake:"{ulid}"`
	ProjectID   string `json:"projectId" validate:"required,max=26" fake:"{ulid}"`
	Name        string `json:"name" validate:"required,max=64" fake:"{sentence:2}"`
	Description string `json:"description,omitempty" validate:"max=255" fake:"{sentence:4}"`
}

type ListCreate struct {
	ProjectID   string `json:"projectId" validate:"required,max=26"`
	Name        string `json:"name" validate:"required,max=64"`
	Description string `json:"description,omitempty" validate:"max=255"`
}

type ListUpdate struct {
	Name        string `json:"name,omitempty" validate:"max=64"`
	Description string `json:"description,omitempty" validate:"max=255"`
}

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

// Project validation methods
func (p *Project) Validate() error {
	return validate.Struct(p)
}

func (pc *ProjectCreate) Validate() error {
	return validate.Struct(pc)
}

func (pu *ProjectUpdate) Validate() error {
	return validate.Struct(pu)
}

// List validation methods
func (l *List) Validate() error {
	return validate.Struct(l)
}

func (lc *ListCreate) Validate() error {
	return validate.Struct(lc)
}

func (lu *ListUpdate) Validate() error {
	return validate.Struct(lu)
}

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

// Random generation methods for Project
func RandomProject() Project {
	var p Project

	err := gofakeit.Struct(&p)
	if err != nil {
		panic(err)
	}

	return p
}

func RandomProjectCreate() ProjectCreate {
	p := RandomProject()

	return ProjectCreate{
		Name:        p.Name,
		Description: p.Description,
	}
}

func RandomProjectUpdate() ProjectUpdate {
	p := RandomProject()

	return ProjectUpdate{
		Name:        p.Name,
		Description: p.Description,
	}
}

// Random generation methods for List
func RandomList() List {
	var l List

	err := gofakeit.Struct(&l)
	if err != nil {
		panic(err)
	}

	return l
}

func RandomListCreate() ListCreate {
	l := RandomList()

	return ListCreate{
		ProjectID:   l.ProjectID,
		Name:        l.Name,
		Description: l.Description,
	}
}

func RandomListUpdate() ListUpdate {
	l := RandomList()

	return ListUpdate{
		Name:        l.Name,
		Description: l.Description,
	}
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
