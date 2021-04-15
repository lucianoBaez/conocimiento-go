package model

import (
	"context"
	"github.com/Fs02/rel"
	"github.com/Fs02/rel/where"
	"time"
)

// User is a model that maps to users table.
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required,min=2,max=50"`
	Username  string    `json:"username" validate:"required,min=2,max=50"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository interfaces for accessing user data
type UserRepository interface {
	FindAll(ctx context.Context, users *[]User) error
	Find(ctx context.Context, user *User, id int64) error
	FindByUsername(ctx context.Context, user *User, username string) error
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
}

// actual implementation
type userRepository struct {
	repository rel.Repository
}

func (ur userRepository) FindAll(ctx context.Context, users *[]User) error {
	return ur.repository.FindAll(ctx, users)
}

func (ur userRepository) Find(ctx context.Context, user *User, id int64) error {
	return ur.repository.Find(ctx, user, where.Eq("id", id))
}

func (ur userRepository) FindByUsername(ctx context.Context, user *User, username string) error {
	return ur.repository.Find(ctx, user, where.Eq("username", username))
}

func (ur userRepository) Create(ctx context.Context, user *User) error {
	return ur.repository.Insert(ctx, user)
}

func (ur userRepository) Update(ctx context.Context, user *User) error {
	return ur.repository.Update(ctx, user)
}

func (ur userRepository) Delete(ctx context.Context, user *User) error {
	return ur.repository.Delete(ctx, user)
}

// NewContactRepository returns a new repository
func NewUserRepository(repo rel.Repository) UserRepository {
	return userRepository{
		repository: repo,
	}
}
