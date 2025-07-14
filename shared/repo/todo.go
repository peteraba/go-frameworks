package repo

import (
	"errors"
	"sync"

	"github.com/peteraba/go-frameworks/shared/model"
)

type TodoRepo interface {
	Create(todo model.Todo) (model.Todo, error)
	GetByID(id string) (model.Todo, error)
	Update(id string, update model.TodoUpdate) (model.Todo, error)
	Delete(id string) error
	List() ([]model.Todo, error)
}

type InMemoryTodoRepo struct {
	mu    sync.RWMutex
	todos map[string]model.Todo
}

func NewInMemoryTodoRepo() *InMemoryTodoRepo {
	return &InMemoryTodoRepo{
		todos: make(map[string]model.Todo),
	}
}

func (r *InMemoryTodoRepo) Create(todo model.Todo) (model.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if todo.ID == "" {
		return model.Todo{}, errors.New("todo ID is required")
	}
	if _, exists := r.todos[todo.ID]; exists {
		return model.Todo{}, errors.New("todo already exists")
	}

	r.todos[todo.ID] = todo

	return todo, nil
}

func (r *InMemoryTodoRepo) GetByID(id string) (model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todo, exists := r.todos[id]
	if !exists {
		return model.Todo{}, errors.New("todo not found")
	}

	return todo, nil
}

func (r *InMemoryTodoRepo) Update(id string, update model.TodoUpdate) (model.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo, exists := r.todos[id]
	if !exists {
		return model.Todo{}, errors.New("todo not found")
	}

	if update.Title != "" {
		todo.Title = update.Title
	}
	if update.Description != "" {
		todo.Description = update.Description
	}
	// Add more fields as needed

	r.todos[id] = todo

	return todo, nil
}

func (r *InMemoryTodoRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[id]; !exists {
		return errors.New("todo not found")
	}

	delete(r.todos, id)

	return nil
}

func (r *InMemoryTodoRepo) List() ([]model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todos := make([]model.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}

	return todos, nil
}
