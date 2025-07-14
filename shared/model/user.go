package model

import "github.com/brianvoe/gofakeit/v7"

type User struct {
	ID    string `json:"id" validate:"required,max=26" fake:"{ulid}"`
	Name  string `json:"name" validate:"required,max=64" fake:"{firstname} {lastname}"`
	Email string `json:"email" validate:"required,email" fake:"{email}"`
}

type UserCreate struct {
	Name  string `json:"name" validate:"required,max=64"`
	Email string `json:"email" validate:"required,email"`
}

type UserUpdate struct {
	Name  string `json:"name,omitempty" validate:"max=64"`
	Email string `json:"email,omitempty" validate:"email"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
func (uc *UserCreate) Validate() error {
	return validate.Struct(uc)
}
func (uu *UserUpdate) Validate() error {
	return validate.Struct(uu)
}

func RandomUser() User {
	var u User
	if err := gofakeit.Struct(&u); err != nil {
		panic(err)
	}
	return u
}

func RandomUserCreate() UserCreate {
	u := RandomUser()
	return UserCreate{Name: u.Name, Email: u.Email}
}
func RandomUserUpdate() UserUpdate {
	u := RandomUser()
	return UserUpdate{Name: u.Name, Email: u.Email}
}
