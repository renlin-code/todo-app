package todo

type TodoCategory struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersCategory struct {
	Id         int
	UserId     int
	CategoryId int
}

type TodoTask struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type CategoriesTask struct {
	Id         int
	CategoryId int
	TaskId     int
}
