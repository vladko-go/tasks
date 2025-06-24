package usersService

import (
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	GetUsers() ([]User, error)
	GetByID(id uint) (User, error)
	Create(UserCreateRequest) (User, error)
	Update(id uint, req UserUpdateRequest) (User, error) // Update(id string, title string, status string) error
	Delete(id uint) error
}

type userService struct {
	repo UserRepository
}

func NewService(r UserRepository) UserService {
	return &userService{repo: r}
}

// Create implements TaskService.
func (s *userService) Create(payload UserCreateRequest) (User, error) {

	user := User{
		ID:        uint(uuid.New().NodeID()[2]),
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return User{}, err
	}

	return user, nil
}

// Delete implements TaskService.
func (s *userService) Delete(id uint) error {

	return s.repo.Delete(id)
}

// GetAll implements TaskService.
func (s *userService) GetUsers() ([]User, error) {
	return s.repo.GetAll()
}

// GetByID implements TaskService.
func (s *userService) GetByID(id uint) (User, error) {

	return s.repo.GetByID(id)
}

// Update implements TaskService.
func (s *userService) Update(id uint, req UserUpdateRequest) (User, error) {

	user, err := s.repo.GetByID(id)

	if err != nil {
		return User{}, err
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Password != "" {
		user.Password = req.Password
	}

	user.UpdatedAt = time.Now()

	if err := s.repo.Update(user); err != nil {
		return User{}, nil
	}

	return user, nil
}
