package repo

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/oklog/ulid/v2"
	"github.com/peteraba/go-frameworks/shared/model"
)

const maxProjectListLength = 100

type ProjectRepo interface {
	Create(project model.ProjectCreate) (model.Project, error)
	GetByID(id string) (model.Project, error)
	Update(id string, update model.ProjectUpdate) (model.Project, error)
	Delete(id string) error
	List() ([]model.Project, error)
}

// ErrProjectNotFound
var ErrProjectNotFound = errors.New("project not found")

type InMemoryProjectRepo struct {
	mu       sync.RWMutex
	projects map[string]model.Project
	ids      []string
	dirty    bool
}

func NewInMemoryProjectRepo() *InMemoryProjectRepo {
	return &InMemoryProjectRepo{
		projects: make(map[string]model.Project),
	}
}

func (r *InMemoryProjectRepo) Create(project model.ProjectCreate) (model.Project, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p := model.Project{
		ID:          ulid.Make().String(),
		Name:        project.Name,
		Description: project.Description,
	}

	if _, exists := r.projects[p.ID]; exists {
		return model.Project{}, fmt.Errorf("not found: %s, err: %w", p.ID, ErrProjectNotFound)
	}

	r.projects[p.ID] = p
	r.ids = append(r.ids, p.ID)
	r.dirty = true

	return p, nil
}

func (r *InMemoryProjectRepo) GetByID(id string) (model.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	project, exists := r.projects[id]
	if !exists {
		return model.Project{}, fmt.Errorf("not found: %s, err: %w", id, ErrProjectNotFound)
	}

	return project, nil
}

func (r *InMemoryProjectRepo) Update(id string, update model.ProjectUpdate) (model.Project, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	project, exists := r.projects[id]
	if !exists {
		return model.Project{}, fmt.Errorf("not found: %s, err: %w", id, ErrProjectNotFound)
	}

	if update.Name != "" {
		project.Name = update.Name
	}
	if update.Description != "" {
		project.Description = update.Description
	}

	r.projects[id] = project

	return project, nil
}

func (r *InMemoryProjectRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.projects[id]; !exists {
		return fmt.Errorf("not found: %s, err: %w", id, ErrProjectNotFound)
	}

	delete(r.projects, id)

	return nil
}

func (r *InMemoryProjectRepo) List() ([]model.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.dirty {
		keys := make([]string, 0, len(r.projects))
		for k := range r.projects {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		r.dirty = false
		r.ids = keys
	}

	l := len(r.projects)
	if l > maxProjectListLength {
		l = maxProjectListLength
	}

	projects := make([]model.Project, 0, l)
	for i, key := range r.ids {
		if i >= l {
			break
		}

		projects = append(projects, r.projects[key])
	}

	return projects, nil
}

func (r *InMemoryProjectRepo) Has(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.projects[id]

	return exists
}
