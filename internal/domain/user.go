package domain

import "time"

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	FindByID(id uint) (User, error)
	GetByUsername(username string) (User, error)
	Create(user User) (User, error)
}
