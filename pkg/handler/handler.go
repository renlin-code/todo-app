package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/renlin-code/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		categories := api.Group("/categories")
		{
			categories.POST("/", h.createCategory)
			categories.GET("/", h.getAllCategories)
			categories.GET("/:id", h.getCategoryById)
			categories.PUT("/:id", h.updateCategory)
			categories.DELETE("/:id", h.deleteCategory)
		}

		task := api.Group("/tasks")
		{
			task.POST("/", h.createTask)
			task.GET("/", h.getAllTasks)
			task.GET("/:id", h.getTaskById)
			task.PUT("/:id", h.updateTask)
			task.DELETE("/:id", h.deleteTask)
		}
	}

	return router
}
