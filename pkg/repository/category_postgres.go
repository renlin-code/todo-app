package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/renlin-code/todo-app"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db}
}

func (r *CategoryPostgres) Create(userId int, category todo.Category) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", categoriesTables)

	row := tx.QueryRow(createCategoryQuery, category.Title, category.Description)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersCategoriesQuery := fmt.Sprintf("INSERT INTO %s (user_id, category_id) VALUES ($1, $2)", usersCategoriesTable)

	_, err = tx.Exec(createUsersCategoriesQuery, userId, id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CategoryPostgres) GetAll(userId int) ([]todo.Category, error) {
	var categories []todo.Category

	query := fmt.Sprintf("SELECT ct.id, ct.title, ct.description FROM %s ct INNER JOIN %s uct ON ct.id = uct.category_id WHERE uct.user_id = $1", categoriesTables, usersCategoriesTable)

	err := r.db.Select(&categories, query, userId)

	return categories, err
}

func (r *CategoryPostgres) GetById(userId, categoryId int) (todo.Category, error) {
	var category todo.Category

	query := fmt.Sprintf("SELECT ct.id, ct.title, ct.description FROM %s ct INNER JOIN %s uct ON ct.id = uct.category_id WHERE uct.user_id = $1 AND uct.category_id = $2", categoriesTables, usersCategoriesTable)

	err := r.db.Get(&category, query, userId, categoryId)

	return category, err
}
