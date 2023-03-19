package database_test

import (
	"context"
	"errors"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type MemoryUserRepository struct {
	Users map[string]*entities.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		Users: make(map[string]*entities.User),
	}
}

func (r *MemoryUserRepository) CreateUser(ctx context.Context, user *entities.User) error {
	r.Users[user.Email] = user
	return nil
}

func (r *MemoryUserRepository) FindByEmail(email string) (*dtos.UserOutputDTO, error) {
	user, ok := r.Users[email]

	if !ok {
		return nil, errors.New("not found")
	}

	return &dtos.UserOutputDTO{
		IDUser:         user.IDUser,
		FullName:       user.FullName,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
	}, nil
}
