package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/renlin-code/todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type Category interface {
	Create(userId int, category todo.Category) (int, error)
	GetAll(userId int) ([]todo.Category, error)
	GetById(userId, categoryId int) (todo.Category, error)
	Update(userId, categoryId int, input todo.UpdateCategoryInput) error
	Delete(userId, categoryId int) error
}

type Task interface {
	Create(categoryId int, task todo.Task) (int, error)
	GetAll(userId, categoryId int) ([]todo.Task, error)
	GetById(userId, taskId int) (todo.Task, error)
	Delete(userId, taskId int) error
	Update(userId, taskId int, input todo.UpdateTaskInput) error
}

type Repository struct {
	Authorization
	Category
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Category:      NewCategoryPostgres(db),
		Task:          NewTaskPostgres(db),
	}
}
