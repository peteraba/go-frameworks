package service

import (
	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
)

type ProjectService struct {
	repo repo.ProjectRepo
}

func NewProjectService(r repo.ProjectRepo) *ProjectService {
	return &ProjectService{repo: r}
}

func (s *ProjectService) Create(pc model.ProjectCreate) (model.Project, error) {
	if err := pc.Validate(); err != nil {
		return model.Project{}, err
	}

	return s.repo.Create(pc)
}

func (s *ProjectService) GetByID(id string) (model.Project, error) {
	return s.repo.GetByID(id)
}

func (s *ProjectService) Update(id string, pu model.ProjectUpdate) (model.Project, error) {
	if err := pu.Validate(); err != nil {
		return model.Project{}, err
	}

	return s.repo.Update(id, pu)
}

func (s *ProjectService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *ProjectService) List() ([]model.Project, error) {
	return s.repo.List()
}
