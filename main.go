package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TaskStatus string

const (
	TODO       TaskStatus = "TODO"
	DONE       TaskStatus = "DONE"
	INPROGRESS TaskStatus = "INPROGRESS"
)

type Task struct {
	ID     string     `json:"id"`
	Title  string     `json: title`
	Status TaskStatus `json:"status"`
}

type TaskCreateRequest struct {
	Title string `json:"title" validate:"required"`
}

type TaskUpdateRequest struct {
	Title  string     `json:"title" validate:"required"`
	Status TaskStatus `json:"status" validate:"oneof=TODO DONE INPROGRESS"`
}

var tasks []Task

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func createTask(c echo.Context) error {
	var req TaskCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	task := Task{
		ID:     uuid.New().String(),
		Title:  req.Title,
		Status: TODO,
	}

	tasks = append(tasks, task)

	return c.JSON(http.StatusCreated, task)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "task " + id + " deleted"})
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "task not found"})
}

func editTask(c echo.Context) error {
	id := c.Param("id")

	var req TaskUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	for i, task := range tasks {
		if task.ID == id {

			if req.Title != "" {
				tasks[i].Title = req.Title
			}
			if req.Status != "" {
				tasks[i].Status = req.Status
			}

			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "task not found"})
}

func main() {

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getTask)
	e.POST("/tasks", createTask)
	e.DELETE("/tasks/:id", deleteTask)
	e.PATCH("/tasks/:id", editTask)

	e.Start("localhost:8080")

}
