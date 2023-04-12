package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/reangeline/workout-plan-go/internal/domain/contracts/usecases"
	"github.com/reangeline/workout-plan-go/internal/dtos"
	"github.com/reangeline/workout-plan-go/internal/validation/protocols"
)

type TrainingController struct {
	trainingUseCase   usecases.TrainingUseCaseInterface
	trainingValidator protocols.TrainingValidatorInterface
}

func NewTrainingController(
	trainingUseCase usecases.TrainingUseCaseInterface,
	trainingValidator protocols.TrainingValidatorInterface,
) *TrainingController {
	return &TrainingController{
		trainingUseCase:   trainingUseCase,
		trainingValidator: trainingValidator,
	}
}

// Create training godoc
// @Summary      Create training
// @Description  Create training
// @Tags         trainings
// @Accept       json
// @Produce      json
// @Param        request     body      dtos.CreateTrainingInput  true  "training request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /trainings [post]
func (u *TrainingController) CreateTraining(w http.ResponseWriter, r *http.Request) {
	var training dtos.CreateTrainingInput
	err := json.NewDecoder(r.Body).Decode(&training)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = u.trainingValidator.ValidateTraining(&training)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = u.trainingUseCase.CreateTraining(ctx, &training)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
