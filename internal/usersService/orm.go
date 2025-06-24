package usersService

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:NULL" json:"deleted_at"`
}

type UserCreateRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserUpdateRequest struct {
	Password  string `json:"password"`
	Email     string `json:"email"`
	UpdatedAt time.Time
}
