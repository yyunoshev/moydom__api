package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"moydom_api/internal/domain"
)

type UserService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// GetUserByID возвращает пользователя по ID
func (s *UserService) GetUserByID(id uint) (domain.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (domain.User, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) CreateUser(input domain.AuthInput) (domain.User, error) {
	existingUser, err := s.userRepo.GetByUsername(input.Username)
	if err != nil {
		return domain.User{}, err
	}
	if existingUser.ID != 0 {
		return domain.User{}, errors.New("username already exists")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{
		Username: input.Username,
		Password: string(passwordHash),
	}
	return s.userRepo.Create(user)
}
