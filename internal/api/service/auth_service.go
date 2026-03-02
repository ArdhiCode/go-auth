package service

import (
	"context"
	"errors"

	"github.com/ArdhiCode/go-auth/internal/api/repository"
	"github.com/ArdhiCode/go-auth/internal/dto"
	"github.com/ArdhiCode/go-auth/internal/entity"
	myjwt "github.com/ArdhiCode/go-auth/internal/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
}

type authService struct {
	userRepo   repository.UserRepository
	jwtService myjwt.JWT
}

func NewAuth(r repository.UserRepository, jwtService myjwt.JWT) AuthService {
	return &authService{
		userRepo:   r,
		jwtService: jwtService,
	}
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.userRepo.Get(ctx, nil, entity.User{Email: req.Email})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.LoginResponse{}, errors.New("user not found")
		}
		return dto.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	token, err := s.jwtService.CreateToken(user.Id.String(), user.Role)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token: token,
	}, nil
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	existing, err := s.userRepo.Get(ctx, nil, entity.User{Email: req.Email})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.RegisterResponse{}, err
	}
	if existing != nil {
		return dto.RegisterResponse{}, errors.New("email already registered")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPass),
		Role:     req.Role,
	}

	result, err := s.userRepo.Create(ctx, nil, user)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		Email: result.Email,
		Name:  result.Name,
		Role:  result.Role,
	}, nil
}
