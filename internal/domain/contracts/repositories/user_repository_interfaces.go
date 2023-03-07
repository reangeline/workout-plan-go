package repositories

import "github.com/reangeline/workout-plan-go/internal/domain/entities"

type UserRepositoryInterface interface {
	CreateUser(user *entities.User) error
}
