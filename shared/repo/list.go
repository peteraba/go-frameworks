package repo

import (
	"errors"
	"sync"

	"github.com/oklog/ulid/v2"
	"github.com/peteraba/go-frameworks/shared/model"
)

type ListRepo interface {
	Create(list model.List) (model.List, error)
	GetByID(id string) (model.List, error)
	Update(id string, update model.ListUpdate) (model.List, error)
	Delete(id string) error
	List() ([]model.List, error)
}

type InMemoryListRepo struct {
	mu    sync.RWMutex
	lists map[string]model.List
}

func NewInMemoryListRepo() *InMemoryListRepo {
	return &InMemoryListRepo{
		lists: make(map[string]model.List),
	}
}

func (r *InMemoryListRepo) Create(list model.ListCreate) (model.List, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	c := model.List{
		ID:          ulid.Make().String(),
		ProjectID:   list.ProjectID,
		Name:        list.Name,
		Description: list.Description,
	}

	if _, exists := r.lists[c.ID]; exists {
		return model.List{}, errors.New("list already exists")
	}

	r.lists[c.ID] = c

	return c, nil
}

func (r *InMemoryListRepo) Has(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.lists[id]

	return exists
}

func (r *InMemoryListRepo) GetByID(id string) (model.List, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list, exists := r.lists[id]
	if !exists {
		return model.List{}, errors.New("list not found")
	}

	return list, nil
}

func (r *InMemoryListRepo) Update(id string, update model.ListUpdate) (model.List, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	list, exists := r.lists[id]
	if !exists {
		return model.List{}, errors.New("list not found")
	}

	if update.Name != "" {
		list.Name = update.Name
	}
	if update.Description != "" {
		list.Description = update.Description
	}

	r.lists[id] = list

	return list, nil
}

func (r *InMemoryListRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.lists[id]; !exists {
		return errors.New("list not found")
	}

	delete(r.lists, id)

	return nil
}

func (r *InMemoryListRepo) List() ([]model.List, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	lists := make([]model.List, 0, len(r.lists))
	for _, list := range r.lists {
		lists = append(lists, list)
	}

	return lists, nil
}
