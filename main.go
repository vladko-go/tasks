package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TaskRequest struct {
	Task string `json:"task"`
}

var task string

func getTask(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, "+task)
}

func createTask(c echo.Context) error {
	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	task = req.Task

	return c.JSON(http.StatusOK, task)
}

func main() {

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getTask)

	e.POST("/tasks", createTask)

	e.Start("localhost:8080")

}
