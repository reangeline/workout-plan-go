package protocols

import "github.com/reangeline/workout-plan-go/internal/dtos"

type UserValidatorInterface interface {
	ValidateUser(user *dtos.UserInputDTO) error
}
