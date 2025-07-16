package repo

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/oklog/ulid/v2"
	"github.com/peteraba/go-frameworks/shared/model"
)

const maxTodoListLength = 1000

type TodoRepo interface {
	Create(todo model.TodoCreate) (model.Todo, error)
	GetByID(id string) (model.Todo, error)
	Update(id string, update model.TodoUpdate) (model.Todo, error)
	Delete(id string) error
	List() ([]model.Todo, error)
}

// ErrTodoNotFound
var ErrTodoNotFound = errors.New("todo item not found")

type InMemoryTodoRepo struct {
	mu    sync.RWMutex
	todos map[string]model.Todo
	ids   []string
	dirty bool
}

func NewInMemoryTodoRepo() *InMemoryTodoRepo {
	return &InMemoryTodoRepo{
		todos: make(map[string]model.Todo),
	}
}

func (r *InMemoryTodoRepo) Create(todo model.TodoCreate) (model.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t := model.Todo{
		ID:          ulid.Make().String(),
		ListID:      todo.ListID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}

	r.todos[t.ID] = t
	r.ids = append(r.ids, t.ID)

	return t, nil
}

func (r *InMemoryTodoRepo) Has(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.todos[id]

	return exists
}

func (r *InMemoryTodoRepo) GetByID(id string) (model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todo, exists := r.todos[id]
	if !exists {
		return model.Todo{}, fmt.Errorf("not found: %s, err: %w", id, ErrTodoNotFound)
	}

	return todo, nil
}

func (r *InMemoryTodoRepo) Update(id string, update model.TodoUpdate) (model.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo, exists := r.todos[id]
	if !exists {
		return model.Todo{}, fmt.Errorf("not found: %s, err: %w", id, ErrTodoNotFound)
	}

	if update.Title != "" {
		todo.Title = update.Title
	}
	if update.Description != "" {
		todo.Description = update.Description
	}
	if update.Completed != nil {
		todo.Completed = *update.Completed
	}

	r.todos[id] = todo

	return todo, nil
}

func (r *InMemoryTodoRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[id]; !exists {
		return fmt.Errorf("not found: %s, err: %w", id, ErrTodoNotFound)
	}

	delete(r.todos, id)

	return nil
}

func (r *InMemoryTodoRepo) List() ([]model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.dirty {
		keys := make([]string, 0, len(r.todos))
		for k := range r.todos {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		r.dirty = false
		r.ids = keys
	}

	l := len(r.todos)
	if l > maxTodoListLength {
		l = maxTodoListLength
	}

	todos := make([]model.Todo, 0, l)
	for i, key := range r.ids {
		if i >= l {
			break
		}

		todos = append(todos, r.todos[key])
	}

	return todos, nil
}
