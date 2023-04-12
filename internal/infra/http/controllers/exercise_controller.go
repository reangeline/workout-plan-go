package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/reangeline/workout-plan-go/internal/domain/contracts/usecases"
	"github.com/reangeline/workout-plan-go/internal/dtos"
	"github.com/reangeline/workout-plan-go/internal/validation/protocols"
)

type ExerciseController struct {
	exerciseUseCase   usecases.ExerciseUseCaseInterface
	exerciseValidator protocols.ExerciseValidatorInterface
}

func NewExerciseController(exerciseUseCase usecases.ExerciseUseCaseInterface, exerciseValidator protocols.ExerciseValidatorInterface) *ExerciseController {
	return &ExerciseController{
		exerciseUseCase:   exerciseUseCase,
		exerciseValidator: exerciseValidator,
	}
}

// Create exercise godoc
// @Summary      Create exercise
// @Description  Create exercise
// @Tags         exercises
// @Accept       json
// @Produce      json
// @Param        request     body      dtos.CreateExerciseInput  true  "exercise request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /exercises [post]
func (u *ExerciseController) CreateExercise(w http.ResponseWriter, r *http.Request) {
	var exercise dtos.CreateExerciseInput
	err := json.NewDecoder(r.Body).Decode(&exercise)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	// err = u.exerciseValidator.ValidateUser(&exercise)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	err = u.exerciseUseCase.CreateExercise(ctx, &exercise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
