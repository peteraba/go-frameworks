package model

import "github.com/brianvoe/gofakeit/v7"

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
