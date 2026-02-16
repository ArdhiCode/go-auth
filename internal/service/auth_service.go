package service

import (
	"errors"

	"github.com/ArdhiCode/go-auth/internal/model"
	"github.com/ArdhiCode/go-auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string
	Email    string
	Password string
	Role     string
}

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

func validateRole(role string) error {
	switch role {
	case RoleUser, RoleAdmin:
		return nil
	default:
		return errors.New("Invalid role")
	}
}

func Register(input RegisterInput) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
		Role:     input.Role,
	}

	return repository.CreateUser(&user)
}
