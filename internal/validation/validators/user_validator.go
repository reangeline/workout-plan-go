package validators

import (
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type UserValidator struct{}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (uv *UserValidator) ValidateUser(user *dtos.UserInputDTO) error {
	// Implementação da validação de usuário aqui

	return nil
}
