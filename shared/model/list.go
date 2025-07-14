package model

import "github.com/brianvoe/gofakeit/v7"

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
