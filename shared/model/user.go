package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID           string   `json:"id" validate:"required,max=26" fake:"{ulid}"`
	Name         string   `json:"name" validate:"required,max=64" fake:"{firstname} {lastname}"`
	Email        string   `json:"email" validate:"required,email" fake:"{email}"`
	Groups       []string `json:"groups" validate:"dive,max=26"`
	PasswordHash []byte   `json:"-"`
	PasswordSalt []byte   `json:"-"`
}

type UserCreate struct {
	Name      string   `json:"name" validate:"required,max=64"`
	Email     string   `json:"email" validate:"required,email"`
	Password  string   `json:"password" validate:"required,min=8"`
	Password2 string   `json:"password2" validate:"required,min=8,eqfield=Password"`
	Groups    []string `json:"groups" validate:"dive,max=26"`
}

type UserUpdate struct {
	Name   string   `json:"name,omitempty" validate:"max=64"`
	Email  string   `json:"email,omitempty" validate:"email"`
	Groups []string `json:"groups,omitempty" validate:"dive,max=26"`
}

type UserLogin struct {
	Name     string `json:"name,omitempty" validate:"max=64"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserPasswordUpdate struct {
	Password  string `json:"password" validate:"required,min=8"`
	Password2 string `json:"password2" validate:"required,min=8,eqfield=Password"`
}

type LoggedInUser struct {
	ID     string
	Name   string
	Groups []string
	jwt.RegisteredClaims
}

// Ensure we implement [jwt.ClaimsValidator] at compile time so we know our custom Validate method is used.
var _ jwt.ClaimsValidator = (*LoggedInUser)(nil)

func (u *User) Validate() error {
	return validate.Struct(u)
}
func (uc *UserCreate) Validate() error {
	return validate.Struct(uc)
}
func (uu *UserUpdate) Validate() error {
	return validate.Struct(uu)
}
func (ul *UserLogin) Validate() error {
	return validate.Struct(ul)
}
func (upu *UserPasswordUpdate) Validate() error {
	return validate.Struct(upu)
}
func (liu *LoggedInUser) Validate() error {
	return validate.Struct(liu)
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
	pw := gofakeit.Password(true, true, true, true, false, 12)
	return UserCreate{
		Name:      u.Name,
		Email:     u.Email,
		Password:  pw,
		Password2: pw,
		Groups:    u.Groups,
	}
}
func RandomUserUpdate() UserUpdate {
	u := RandomUser()

	return UserUpdate{Name: u.Name, Email: u.Email, Groups: u.Groups}
}
func RandomUserLogin() UserLogin {
	u := RandomUser()

	p1 := gofakeit.Password(true, true, true, true, true, 12)

	return UserLogin{Name: u.Name, Password: p1}
}
func RandomUserPasswordUpdate() UserPasswordUpdate {
	p1 := gofakeit.Password(true, true, true, true, true, 12)

	return UserPasswordUpdate{Password: p1, Password2: p1}
}
