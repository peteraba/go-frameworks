package service

import (
	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
)

type ListService struct {
	repo repo.ListRepo
}

func NewListService(r repo.ListRepo) *ListService {
	return &ListService{repo: r}
}

func (s *ListService) Create(lc model.ListCreate) (model.List, error) {
	if err := lc.Validate(); err != nil {
		return model.List{}, err
	}

	return s.repo.Create(lc)
}

func (s *ListService) GetByID(id string) (model.List, error) {
	return s.repo.GetByID(id)
}

func (s *ListService) Update(id string, lu model.ListUpdate) (model.List, error) {
	if err := lu.Validate(); err != nil {
		return model.List{}, err
	}

	return s.repo.Update(id, lu)
}

func (s *ListService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *ListService) List() ([]model.List, error) {
	return s.repo.List()
}
