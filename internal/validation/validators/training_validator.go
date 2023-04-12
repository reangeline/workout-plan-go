package validators

import (
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type TrainingValidator struct{}

func NewTrainingValidator() *TrainingValidator {
	return &TrainingValidator{}
}

func (uv *TrainingValidator) ValidateTraining(exercise *dtos.CreateTrainingInput) error {

	return nil
}
