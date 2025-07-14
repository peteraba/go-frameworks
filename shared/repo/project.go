package repo

import (
	"errors"
	"sync"

	"github.com/peteraba/go-frameworks/shared/model"
)

type ProjectRepo interface {
	Create(project model.Project) (model.Project, error)
	GetByID(id string) (model.Project, error)
	Update(id string, update model.ProjectUpdate) (model.Project, error)
	Delete(id string) error
	List() ([]model.Project, error)
	Has(id string) bool
}

type InMemoryProjectRepo struct {
	mu       sync.RWMutex
	projects map[string]model.Project
}

func NewInMemoryProjectRepo() *InMemoryProjectRepo {
	return &InMemoryProjectRepo{
		projects: make(map[string]model.Project),
	}
}

func (r *InMemoryProjectRepo) Create(project model.Project) (model.Project, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if project.ID == "" {
		return model.Project{}, errors.New("project ID is required")
	}
	if _, exists := r.projects[project.ID]; exists {
		return model.Project{}, errors.New("project already exists")
	}

	r.projects[project.ID] = project

	return project, nil
}

func (r *InMemoryProjectRepo) GetByID(id string) (model.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	project, exists := r.projects[id]
	if !exists {
		return model.Project{}, errors.New("project not found")
	}

	return project, nil
}

func (r *InMemoryProjectRepo) Update(id string, update model.ProjectUpdate) (model.Project, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	project, exists := r.projects[id]
	if !exists {
		return model.Project{}, errors.New("project not found")
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
		return errors.New("project not found")
	}

	delete(r.projects, id)

	return nil
}

func (r *InMemoryProjectRepo) List() ([]model.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	projects := make([]model.Project, 0, len(r.projects))
	for _, project := range r.projects {
		projects = append(projects, project)
	}

	return projects, nil
}
