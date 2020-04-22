package main

import (
	"github.com/Lucku/go-workshop/rest-service/todo"
	"log"

	"github.com/gin-gonic/gin"
)

func start() error {

	repo := todo.NewMapRepository()
	api := todo.NewAPI(repo)

	router := gin.Default()

	router.GET("/todos/", api.GetAllTodos)
	router.GET("/todos/:id", api.GetTodoByID)
	router.POST("/todos/", api.CreateTodo)
	router.DELETE("/todos/:id", api.DeleteTodo)

	if err := router.Run(":8080"); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := start(); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
