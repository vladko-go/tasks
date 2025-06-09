package main

import (
	"pet-project/internal/db"
	"pet-project/internal/handlers"
	"pet-project/internal/taskService"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	taskRepository := taskService.NewTaskRepository(db)

	taskService := taskService.NewTaskService(taskRepository)

	handler := handlers.NewTaskHandler(taskService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", handler.GetTasks)
	e.POST("/tasks", handler.CreateTask)
	e.DELETE("/tasks/:id", handler.DeleteTask)
	e.PATCH("/tasks/:id", handler.EditTask)

	e.Start("localhost:8080")

}
