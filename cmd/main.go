package main

import (
	"log"

	"github.com/renlin-code/todo-app"
	"github.com/renlin-code/todo-app/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)

	err := srv.Run("8080", handler.InitRoutes())
	if err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
