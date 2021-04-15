package service

import (
	"context"
	"fmt"
	"github.com/d-Una-Interviews/svc_aut/pkg/model"
	"github.com/d-Una-Interviews/svc_aut/pkg/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

var (
	userServiceLogger, _ = zap.NewProduction(zap.Fields(zap.String("type", "service")))
	userRepository       = repository.InitRepository()
)

// UserService Define services for User Management
type UserService interface {
	// Create a User
	CreateUser(ctx context.Context, request *model.User) (*model.User, error)
	FindAll(ctx context.Context, users *[]model.User) error
	FindUsername(ctx context.Context, user *model.User, username string) error
}

type userService struct {
	Repo      model.UserRepository
	Validator *validator.Validate
}

func (s *userService) FindUsername(ctx context.Context, user *model.User, username string) error {

	err := s.Repo.FindByUsername(ctx, user, username)
	if err != nil {
		userServiceLogger.Error(err.Error(), zap.Error(err))
	}
	return err
}

func (s *userService) FindAll(ctx context.Context, users *[]model.User) error {
	err := s.Repo.FindAll(ctx, users)
	if err != nil {
		userServiceLogger.Error(err.Error(), zap.Error(err))
	}
	return err
}

func (s *userService) CreateUser(ctx context.Context, request *model.User) (*model.User, error) {
	// Validate the user
	err := s.Validator.Struct(request)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
		userServiceLogger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	password := []byte(request.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		userServiceLogger.Error(err.Error(), zap.Error(err))
		return nil, err
	}
	fmt.Println(string(hashedPassword))

	request.Password = string(hashedPassword)

	err = s.Repo.Create(ctx, request)
	if err != nil {
		userServiceLogger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	return request, nil
}

// NewUserService Create a new User Service
func NewUserService() UserService {
	return &userService{
		Repo:      model.NewUserRepository(userRepository),
		Validator: validator.New(),
	}
}
