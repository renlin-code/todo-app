package service

import (
	"github.com/renlin-code/todo-app"
	"github.com/renlin-code/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Category interface {
}

type Task interface {
}

type Service struct {
	Authorization
	Category
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
	}
}
