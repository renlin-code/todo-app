package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/renlin-code/todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type Category interface {
}

type Task interface {
}

type Repository struct {
	Authorization
	Category
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
