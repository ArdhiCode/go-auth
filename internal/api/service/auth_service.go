package service

import (
	"errors"
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

// func Register(input RegisterInput) error {
// 	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

// 	user := entity.User{
// 		Name:     input.Name,
// 		Email:    input.Email,
// 		Password: string(hashed),
// 		Role:     input.Role,
// 	}

// 	return repository.CreateUser(&user)
// }
