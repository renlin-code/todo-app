package service

import (
	"github.com/renlin-code/todo-app"
	"github.com/renlin-code/todo-app/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) Create(userId int, category todo.Category) (int, error) {
	return s.repo.Create(userId, category)
}

func (s *CategoryService) GetAll(userId int) ([]todo.Category, error) {
	return s.repo.GetAll(userId)
}

func (s *CategoryService) GetById(userId, categoryId int) (todo.Category, error) {
	return s.repo.GetById(userId, categoryId)
}
