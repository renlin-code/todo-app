package service

import (
	"github.com/renlin-code/todo-app"
	"github.com/renlin-code/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Category interface {
	Create(userId int, category todo.Category) (int, error)
	GetAll(userId int) ([]todo.Category, error)
	GetById(userId, categoryId int) (todo.Category, error)
	Update(userId, categoryId int, input todo.UpdateCategoryInput) error
	Delete(userId, categoryId int) error
}

type Task interface {
	Create(userId, categoryId int, task todo.Task) (int, error)
	GetAll(userId, categoryId int) ([]todo.Task, error)
	GetById(userId, taskId int) (todo.Task, error)
	Delete(userId, taskId int) error
	Update(userId, taskId int, input todo.UpdateTaskInput) error
}

type Service struct {
	Authorization
	Category
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		Category:      NewCategoryService(repos.Category),
		Task:          NewTaskService(repos.Task, repos.Category),
	}
}
