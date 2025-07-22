package repo

import (
	"errors"
	"sort"
	"sync"

	"github.com/oklog/ulid/v2"
	"github.com/peteraba/go-frameworks/shared/model"
)

const maxListListLength = 100

type ListRepo interface {
	Create(list model.ListCreate) (model.List, error)
	GetByID(id string) (model.List, error)
	Update(id string, update model.ListUpdate) (model.List, error)
	Delete(id string) error
	List() ([]model.List, error)
}

type InMemoryListRepo struct {
	mu    sync.RWMutex
	lists map[string]model.List
	ids   []string
	dirty bool
}

func NewInMemoryListRepo() *InMemoryListRepo {
	return &InMemoryListRepo{
		lists: make(map[string]model.List),
	}
}

func (r *InMemoryListRepo) Create(list model.ListCreate) (model.List, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	listModel := model.List{
		ID:          ulid.Make().String(),
		ProjectID:   list.ProjectID,
		Name:        list.Name,
		Description: list.Description,
	}

	if _, exists := r.lists[listModel.ID]; exists {
		return model.List{}, errors.New("list already exists")
	}

	r.lists[listModel.ID] = listModel
	r.ids = append(r.ids, listModel.ID)
	r.dirty = false

	return listModel, nil
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

	if r.dirty {
		keys := make([]string, 0, len(r.lists))
		for k := range r.lists {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		r.dirty = false
		r.ids = keys
	}

	l := len(r.lists)
	if l > maxListListLength {
		l = maxListListLength
	}

	lists := make([]model.List, 0, l)
	for i, key := range r.ids {
		if i >= l {
			break
		}

		lists = append(lists, r.lists[key])
	}

	return lists, nil
}
