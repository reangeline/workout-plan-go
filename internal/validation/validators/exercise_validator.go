package validators

import (
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type ExerciseValidator struct{}

func NewExerciseValidator() *ExerciseValidator {
	return &ExerciseValidator{}
}

func (uv *ExerciseValidator) ValidateExercise(exercise *dtos.CreateExerciseInput) error {

	return nil
}
