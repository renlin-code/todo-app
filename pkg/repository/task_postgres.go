package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/renlin-code/todo-app"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db}
}

func (r *TaskPostgres) Create(categoryId int, task todo.Task) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var taskId int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", tasksTable)

	row := tx.QueryRow(createTaskQuery, task.Title, task.Description)

	if err := row.Scan(&taskId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createCategoriesTasksQuery := fmt.Sprintf("INSERT INTO %s (category_id, task_id) VALUES ($1, $2)", categoriesTasksTables)

	_, err = tx.Exec(createCategoriesTasksQuery, categoryId, taskId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return taskId, tx.Commit()
}

func (r *TaskPostgres) GetAll(userId, categoryId int) ([]todo.Task, error) {
	var tasks []todo.Task

	query := fmt.Sprintf("SELECT tt.id, tt.title, tt.description, tt.done FROM %s tt INNER JOIN %s ctt ON tt.id = ctt.task_id INNER JOIN %s uct ON ctt.category_id = uct.category_id WHERE uct.user_id = $1 AND ctt.category_id = $2", tasksTable, categoriesTasksTables, usersCategoriesTable)

	err := r.db.Select(&tasks, query, userId, categoryId)

	return tasks, err
}

func (r *TaskPostgres) GetById(userId, taskId int) (todo.Task, error) {
	var task todo.Task

	query := fmt.Sprintf("SELECT tt.id, tt.title, tt.description, tt.done FROM %s tt INNER JOIN %s ctt ON ctt.task_id = tt.id INNER JOIN %s uct ON uct.category_id = ctt.category_id WHERE uct.user_id = $1 AND tt.id = $2", tasksTable, categoriesTasksTables, usersCategoriesTable)

	err := r.db.Get(&task, query, userId, taskId)

	return task, err
}

func (r *TaskPostgres) Update(userId, taskId int, input todo.UpdateTaskInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	var id int
	query := fmt.Sprintf("UPDATE %s tt SET %s FROM %s ctt, %s uct WHERE tt.id = ctt.task_id AND ctt.category_id = uct.category_id AND uct.user_id = $%d AND tt.id = $%d RETURNING tt.id", tasksTable, setQuery, categoriesTasksTables, usersCategoriesTable, argId, argId+1)

	args = append(args, userId, taskId)

	row := r.db.QueryRow(query, args...)

	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (r *TaskPostgres) Delete(userId, taskId int) error {
	var id int
	query := fmt.Sprintf("DELETE FROM %s tt USING %s ctt, %s uct WHERE tt.id = ctt.task_id AND ctt.category_id = uct.category_id AND uct.user_id = $1 AND tt.id = $2 RETURNING tt.id", tasksTable, categoriesTasksTables, usersCategoriesTable)

	row := r.db.QueryRow(query, userId, taskId)
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}
