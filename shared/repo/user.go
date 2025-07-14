package repo

import (
	"errors"
	"sync"

	"github.com/oklog/ulid/v2"
	"github.com/peteraba/go-frameworks/shared/model"
)

type UserRepo interface {
	Create(user model.UserCreate) (model.User, error)
	GetByID(id string) (model.User, error)
	Update(id string, update model.UserUpdate) (model.User, error)
	Delete(id string) error
	List() ([]model.User, error)
}

type InMemoryUserRepo struct {
	mu    sync.RWMutex
	users map[string]model.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]model.User),
	}
}

func (r *InMemoryUserRepo) Create(user model.UserCreate) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	u := model.User{
		ID:    ulid.Make().String(),
		Name:  user.Name,
		Email: user.Email,
	}

	if _, exists := r.users[u.ID]; exists {
		return model.User{}, errors.New("user already exists")
	}

	r.users[u.ID] = u

	return u, nil
}

func (r *InMemoryUserRepo) GetByID(id string) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *InMemoryUserRepo) Update(id string, update model.UserUpdate) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return model.User{}, errors.New("user not found")
	}

	if update.Name != "" {
		user.Name = update.Name
	}
	if update.Email != "" {
		user.Email = update.Email
	}

	r.users[id] = user

	return user, nil
}

func (r *InMemoryUserRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)

	return nil
}

func (r *InMemoryUserRepo) List() ([]model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]model.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

func (r *InMemoryUserRepo) Has(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.users[id]

	return exists
}
