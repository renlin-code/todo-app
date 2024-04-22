package service

import "github.com/renlin-code/todo-app/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
