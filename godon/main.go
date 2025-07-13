package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/abemedia/go-don"
	_ "github.com/abemedia/go-don/encoding/json" // Enable JSON parsing & rendering.

	"github.com/peteraba/go-frameworks/shared"
)

func CreateTodo(ctx context.Context, req shared.TodoCreate) (*shared.Todo, error) {
	if req.Title == "" {
		return nil, don.Error(errors.New("missing title"), http.StatusBadRequest)
	}

	res := &shared.Todo{
		ID:          "01K02MFYHQTCZRPAZGMYC1CZFH",
		Title:       "This is a title",
		Description: "This is a description",
		Complete:    false,
	}

	return res, nil
}

func Pong(context.Context, any) (string, error) {
	return "pong", nil
}

func main() {
	r := don.New(&don.Config{
		DefaultEncoding: "application/json",
	})
	r.Get("/ping", don.H(Pong)) // Handlers are wrapped with `don.H`.
	r.Post("/todos", don.H(CreateTodo))
	r.Put("/todos/:id", don.H(CreateTodo))
	r.ListenAndServe(":8080")
}
