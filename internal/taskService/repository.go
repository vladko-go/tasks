package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id string) (Task, error)
	Create(task Task) error
	Update(task Task) error
	Delete(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error

	return tasks, err

}

func (r *taskRepository) GetByID(id string) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskRepository) Create(task Task) error {
	return r.db.Create(&task).Error
}

func (r *taskRepository) Update(task Task) error {
	return r.db.Save(&task).Error
}

func (r *taskRepository) Delete(id string) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
