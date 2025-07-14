package repo

import (
	"errors"
	"sync"

	"github.com/oklog/ulid/v2"
	"github.com/peteraba/go-frameworks/shared/model"
)

type AuthGroupRepo interface {
	Create(group model.AuthGroupCreate) (model.AuthGroup, error)
	GetByID(id string) (model.AuthGroup, error)
	Update(id string, update model.AuthGroupUpdate) (model.AuthGroup, error)
	Delete(id string) error
	List() ([]model.AuthGroup, error)
}

type InMemoryAuthGroupRepo struct {
	mu     sync.RWMutex
	groups map[string]model.AuthGroup
}

func NewInMemoryAuthGroupRepo() *InMemoryAuthGroupRepo {
	return &InMemoryAuthGroupRepo{
		groups: make(map[string]model.AuthGroup),
	}
}

func (r *InMemoryAuthGroupRepo) Create(group model.AuthGroupCreate) (model.AuthGroup, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	a := model.AuthGroup{
		ID:    ulid.Make().String(),
		Name:  group.Name,
		Users: group.Users,
	}

	if _, exists := r.groups[a.ID]; exists {
		return model.AuthGroup{}, errors.New("auth group already exists")
	}

	r.groups[a.ID] = a

	return a, nil
}

func (r *InMemoryAuthGroupRepo) GetByID(id string) (model.AuthGroup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	group, exists := r.groups[id]
	if !exists {
		return model.AuthGroup{}, errors.New("auth group not found")
	}

	return group, nil
}

func (r *InMemoryAuthGroupRepo) Update(id string, update model.AuthGroupUpdate) (model.AuthGroup, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	group, exists := r.groups[id]
	if !exists {
		return model.AuthGroup{}, errors.New("auth group not found")
	}

	if update.Name != "" {
		group.Name = update.Name
	}
	if len(update.Users) > 0 {
		group.Users = update.Users
	}

	r.groups[id] = group

	return group, nil
}

func (r *InMemoryAuthGroupRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.groups[id]; !exists {
		return errors.New("auth group not found")
	}

	delete(r.groups, id)

	return nil
}

func (r *InMemoryAuthGroupRepo) List() ([]model.AuthGroup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groups := make([]model.AuthGroup, 0, len(r.groups))
	for _, group := range r.groups {
		groups = append(groups, group)
	}

	return groups, nil
}
