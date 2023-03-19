package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/reangeline/workout-plan-go/internal/domain/contracts/usecases"
	"github.com/reangeline/workout-plan-go/internal/dtos"
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
	var user dtos.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err = u.userUseCase.CreateUser(ctx, &user)
	if err != nil {
		http.Error(w, "erro ao criar usuario", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
