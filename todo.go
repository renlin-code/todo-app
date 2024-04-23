package todo

type Category struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersCategories struct {
	Id         int
	UserId     int
	CategoryId int
}

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type CategoriesTasks struct {
	Id         int
	CategoryId int
	TaskId     int
}
