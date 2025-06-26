package main

import (
	"log"
	"pet-project/internal/database"
	"pet-project/internal/handlers"
	"pet-project/internal/tasksService"
	"pet-project/internal/usersService"
	"pet-project/internal/web/tasks"
	"pet-project/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}
	err = database.DB.AutoMigrate(&tasksService.Task{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}

	taskRepo := tasksService.NewTaskRepository(database.DB)
	taskService := tasksService.NewService(taskRepo)

	taskHandler := handlers.NewTaskHandler(taskService)

	userRepo := usersService.NewUserRepository(database.DB)
	userService := usersService.NewService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTasksHandler := tasks.NewStrictHandler(taskHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictTasksHandler)

	strictUsersHandler := users.NewStrictHandler(userHandler, nil) // тут будет ошибка
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
