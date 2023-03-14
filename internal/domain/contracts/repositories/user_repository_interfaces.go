package repositories

import (
	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type UserRepositoryInterface interface {
	CreateUser(user *entities.User) error
	FindByEmail(email string) (dtos.UserOutputDTO, error)
}
