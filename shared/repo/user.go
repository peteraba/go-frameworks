package repo

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/oklog/ulid/v2"
	"github.com/peteraba/go-frameworks/shared/model"
)

const maxUserListLength = 100

type UserRepo interface {
	Create(user model.UserCreate, passwordHash, passwordSalt []byte) (model.User, error)
	GetByID(id string) (model.User, error)
	Update(id string, update model.UserUpdate) (model.User, error)
	Delete(id string) error
	List() ([]model.User, error)
}

// ErrUserNotFound
var ErrUserNotFound = errors.New("user not found")

type InMemoryUserRepo struct {
	mu    sync.RWMutex
	users map[string]model.User
	keys  []string
	dirty bool
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]model.User),
	}
}

func (r *InMemoryUserRepo) Create(uc model.UserCreate, passwordHash, passwordSalt []byte) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	u := model.User{
		ID:           ulid.Make().String(),
		Name:         uc.Name,
		Email:        uc.Email,
		Groups:       uc.Groups,
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
	}

	u.ID = ulid.Make().String()

	r.users[u.ID] = u
	r.keys = append(r.keys, u.ID)
	r.dirty = true

	return u, nil
}

func (r *InMemoryUserRepo) GetByID(id string) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return model.User{}, fmt.Errorf("not found: %s, err: %w", id, ErrUserNotFound)
	}

	return user, nil
}

func (r *InMemoryUserRepo) Update(id string, update model.UserUpdate) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return model.User{}, fmt.Errorf("not found: %s, err: %w", id, ErrUserNotFound)
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

func (r *InMemoryUserRepo) UpdatePassword(id string, passwordHash []byte) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return model.User{}, fmt.Errorf("not found: %s, err: %w", id, ErrUserNotFound)
	}

	user.PasswordHash = passwordHash

	r.users[id] = user

	return user, nil
}

func (r *InMemoryUserRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return fmt.Errorf("not found: %s, err: %w", id, ErrUserNotFound)
	}

	delete(r.users, id)

	return nil
}

func (r *InMemoryUserRepo) List() ([]model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.dirty {
		keys := make([]string, 0, len(r.users))
		for k := range r.users {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		r.dirty = false
		r.keys = keys
	}

	l := len(r.users)
	if l > maxUserListLength {
		l = maxUserListLength
	}

	users := make([]model.User, 0, l)
	for i, key := range r.keys {
		if i >= l {
			break
		}

		users = append(users, r.users[key])
	}

	return users, nil
}

func (r *InMemoryUserRepo) Has(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.users[id]

	return exists
}
