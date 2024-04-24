package service

import (
	"github.com/renlin-code/todo-app"
	"github.com/renlin-code/todo-app/pkg/repository"
)

type TaskService struct {
	repo         repository.Task
	categoryRepo repository.Category
}

func NewTaskService(repo repository.Task, categoryRepo repository.Category) *TaskService {
	return &TaskService{repo, categoryRepo}
}

func (s *TaskService) Create(userId, categoryId int, task todo.Task) (int, error) {
	_, err := s.categoryRepo.GetById(userId, categoryId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(categoryId, task)
}

func (s *TaskService) GetAll(userId, categoryId int) ([]todo.Task, error) {
	_, err := s.categoryRepo.GetById(userId, categoryId)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(userId, categoryId)
}

func (s *TaskService) GetById(userId, taskId int) (todo.Task, error) {
	return s.repo.GetById(userId, taskId)
}

func (s *TaskService) Update(userId, taskId int, input todo.UpdateTaskInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, taskId, input)
}

func (s *TaskService) Delete(userId, taskId int) error {
	return s.repo.Delete(userId, taskId)
}
