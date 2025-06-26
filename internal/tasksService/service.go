package tasksService

import (
	"github.com/google/uuid"
)

type TaskService interface {
	GetTasks() ([]Task, error)
	GetByID(id uint) (Task, error)
	Create(Task) (Task, error)
	Update(id uint, req TaskUpdateRequest) (Task, error) // Update(id string, title string, status string) error
	Delete(id uint) error
}

type taskService struct {
	repo TaskRepository
}

func NewService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

// Create implements TaskService.
func (t *taskService) Create(payload Task) (Task, error) {

	task := Task{
		ID:     uint(uuid.New().NodeID()[2]),
		Task:   payload.Task,
		IsDone: payload.IsDone,
		UserId: payload.UserId,
	}

	if err := t.repo.Create(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

// Delete implements TaskService.
func (t *taskService) Delete(id uint) error {

	return t.repo.Delete(id)
}

// GetAll implements TaskService.
func (t *taskService) GetTasks() ([]Task, error) {
	return t.repo.GetAll()
}

// GetByID implements TaskService.
func (t *taskService) GetByID(id uint) (Task, error) {

	return t.repo.GetByID(id)
}

// Update implements TaskService.
func (t *taskService) Update(id uint, req TaskUpdateRequest) (Task, error) {

	task, err := t.repo.GetByID(id)

	if err != nil {
		return Task{}, err
	}

	if req.Task != "" {
		task.Task = req.Task
	}

	if req.IsDone {
		task.IsDone = req.IsDone
	}

	if err := t.repo.Update(task); err != nil {
		return Task{}, nil
	}

	return task, nil
}
