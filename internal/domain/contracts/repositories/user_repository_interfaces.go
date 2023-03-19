package repositories

import (
	"context"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *entities.User) error
	FindByEmail(email string) (*dtos.UserOutputDTO, error)
}
