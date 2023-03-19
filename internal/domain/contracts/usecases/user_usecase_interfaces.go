package usecases

import (
	"context"

	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, input *dtos.CreateUserInput) error
	CheckEmailExists(email string) (bool, error)
}
