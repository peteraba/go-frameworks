package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/oklog/ulid/v2"
)

func InitFaker() {
	gofakeit.AddFuncLookup("ulid", gofakeit.Info{
		Category:    "custom",
		Description: "ULID",
		Example:     "01K02G50QEGXXK5ZBGHDNPKYHY",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			return ulid.Make(), nil
		},
	})
}
