package controllers

import (
	"net/http"

	"github.com/reangeline/workout-plan-go/internal/domain/contracts/usecases"
	"github.com/reangeline/workout-plan-go/internal/validation/protocols"
)

type UserController struct {
	userUseCase   usecases.UserUseCaseInterface
	userValidator protocols.UserValidatorInterface
}

func NewUserController(userUseCase usecases.UserUseCaseInterface, userValidator protocols.UserValidatorInterface) *UserController {
	return &UserController{
		userUseCase:   userUseCase,
		userValidator: userValidator,
	}
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

}
