package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/oklog/ulid/v2"
)

func NewFaker() *gofakeit.Faker {
	gofakeit.AddFuncLookup("friendname", gofakeit.Info{
		Category:    "custom",
		Description: "ULID",
		Example:     "01K02G50QEGXXK5ZBGHDNPKYHY",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			return ulid.Make(), nil
		},
	})

	faker := gofakeit.New(0)

	return faker
}
