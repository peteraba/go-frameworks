package model

import "github.com/brianvoe/gofakeit/v7"

type AuthGroup struct {
	ID   string `json:"id" validate:"required,max=26" fake:"{ulid}"`
	Name string `json:"name" validate:"required,max=64" fake:"{word}"`
}

type AuthGroupCreate struct {
	Name string `json:"name" validate:"required,max=64"`
}

type AuthGroupUpdate struct {
	Name string `json:"name,omitempty" validate:"max=64"`
}

func (a *AuthGroup) Validate() error {
	return validate.Struct(a)
}
func (ac *AuthGroupCreate) Validate() error {
	return validate.Struct(ac)
}
func (au *AuthGroupUpdate) Validate() error {
	return validate.Struct(au)
}

func RandomAuthGroup() AuthGroup {
	var a AuthGroup
	if err := gofakeit.Struct(&a); err != nil {
		panic(err)
	}
	return a
}
func RandomAuthGroupCreate() AuthGroupCreate {
	a := RandomAuthGroup()
	return AuthGroupCreate{Name: a.Name}
}
func RandomAuthGroupUpdate() AuthGroupUpdate {
	a := RandomAuthGroup()
	return AuthGroupUpdate{Name: a.Name}
}
