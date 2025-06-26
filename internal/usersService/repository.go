package usersService

import (
	taskService "pet-project/internal/tasksService"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id uint) (User, error)
	Create(user User) error
	Update(user User) error
	Delete(id uint) error
	GetTasksByUserId(id uint) ([]taskService.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error

	return users, err

}

func (r *userRepository) GetByID(id uint) (User, error) {
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) Create(user User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) Update(user User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}

func (r *userRepository) GetTasksByUserId(id uint) ([]taskService.Task, error) {
	var tasks []taskService.Task
	err := r.db.Where("user_id = ?", id).Find(&tasks).Error
	return tasks, err
}
