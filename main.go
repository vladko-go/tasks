package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	TODO       TaskStatus = "TODO"
	DONE       TaskStatus = "DONE"
	INPROGRESS TaskStatus = "INPROGRESS"
)

type Task struct {
	ID     string     `gorm: "primaryKey" json:"id"`
	Title  string     `json:title"`
	Status TaskStatus `json:"status"`
}

type TaskCreateRequest struct {
	Title string `json:"title" validate:"required"`
}

type TaskUpdateRequest struct {
	Title  string     `json:"title" validate:"required"`
	Status TaskStatus `json:"status" validate:"oneof=TODO DONE INPROGRESS"`
}

var db *gorm.DB

func initDB() {
	dns := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}

	if err = db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

var tasks []Task

func getTask(c echo.Context) error {
	var tasks []Task

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
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

	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, task)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Task{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}
	return c.NoContent(http.StatusNoContent)
}

func editTask(c echo.Context) error {
	id := c.Param("id")

	var req TaskUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var task Task

	if err := db.First(&task, "id =?", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "task not found"})
	}

	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Status != "" {
		task.Status = req.Status
	}

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}
	return c.JSON(http.StatusOK, task)

}

func main() {
	initDB()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getTask)
	e.POST("/tasks", createTask)
	e.DELETE("/tasks/:id", deleteTask)
	e.PATCH("/tasks/:id", editTask)

	e.Start("localhost:8080")

}
