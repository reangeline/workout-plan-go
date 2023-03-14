package usecases

import "github.com/reangeline/workout-plan-go/internal/dtos"

type UserUseCaseInterface interface {
	CreateUser(input *dtos.UserInputDTO) error
	CheckEmailExists(email string) (bool, error)
}
