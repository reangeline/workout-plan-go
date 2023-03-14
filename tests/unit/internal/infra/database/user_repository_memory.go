package database_test

import (
	"errors"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type MemoryUserRepository struct {
	users map[string]*entities.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[string]*entities.User),
	}
}

func (r *MemoryUserRepository) CreateUser(user *entities.User) error {
	r.users[user.Email] = user
	return nil
}

func (r *MemoryUserRepository) FindByEmail(email string) (dtos.UserOutputDTO, error) {
	user, ok := r.users[email]

	if !ok {
		return dtos.UserOutputDTO{}, errors.New("not found")
	}

	output := dtos.UserOutputDTO{
		ID:             user.ID,
		FullName:       user.FullName,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
	}

	return output, nil
}
