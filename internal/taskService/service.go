package taskService

import (
	"errors"

	"github.com/google/uuid"
)

type TaskService interface {
	GetAll() ([]Task, error)
	GetByID(id string) (Task, error)
	Create(title string) (Task, error)
	Update(id string, req TaskUpdateRequest) (Task, error) // Update(id string, title string, status string) error
	Delete(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

// Create implements TaskService.
func (t *taskService) Create(title string) (Task, error) {

	task := Task{
		ID:     uuid.New().String(),
		Title:  title,
		Status: TODO,
	}

	if err := t.repo.Create(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

// Delete implements TaskService.
func (t *taskService) Delete(id string) error {
	if len(id) < 8 {
		return errors.New("invalid id")
	}
	return t.repo.Delete(id)
}

// GetAll implements TaskService.
func (t *taskService) GetAll() ([]Task, error) {
	return t.repo.GetAll()
}

// GetByID implements TaskService.
func (t *taskService) GetByID(id string) (Task, error) {
	if len(id) < 8 {
		return Task{}, errors.New("invalid id")
	}
	return t.repo.GetByID(id)
}

// Update implements TaskService.
func (t *taskService) Update(id string, req TaskUpdateRequest) (Task, error) {

	if len(id) < 8 {
		return Task{}, errors.New("invalid id")
	}

	task, err := t.repo.GetByID(id)

	if err != nil {
		return Task{}, err
	}

	if req.Title != "" {
		task.Title = req.Title
	}

	if req.Status != "" {
		task.Status = req.Status
	}

	if err := t.repo.Update(task); err != nil {
		return Task{}, nil
	}

	return task, nil
}
