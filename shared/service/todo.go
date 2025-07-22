package service

import (
	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
)

type TodoService struct {
	repo repo.TodoRepo
}

func NewTodoService(r repo.TodoRepo) *TodoService {
	return &TodoService{repo: r}
}

func (s *TodoService) Create(tc model.TodoCreate) (model.Todo, error) {
	if err := tc.Validate(); err != nil {
		return model.Todo{}, err
	}

	return s.repo.Create(tc)
}

func (s *TodoService) GetByID(id string) (model.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *TodoService) Update(id string, tu model.TodoUpdate) (model.Todo, error) {
	if err := tu.Validate(); err != nil {
		return model.Todo{}, err
	}

	return s.repo.Update(id, tu)
}

func (s *TodoService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *TodoService) List() ([]model.Todo, error) {
	return s.repo.List()
}
