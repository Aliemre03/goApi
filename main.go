package main

import (
	"github.com/gofiber/fiber/v2"
	"learnApi/app"
	"learnApi/config"
	"learnApi/repository"
	"learnApi/service"
)

func main() {
	appRoute := fiber.New()
	config.ConnectDB()
	dbClient := config.GetCollection(config.DB, "todos")

	TodoRepositoryDb := repository.NewTodoRepositoryDB(dbClient)

	td := app.TodoHandler{Service: service.NewTodoService(TodoRepositoryDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todo", td.GetAllTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Listen("localhost:8081")
}
