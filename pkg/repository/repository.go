package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
