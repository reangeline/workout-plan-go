package protocols

import "github.com/reangeline/workout-plan-go/internal/dtos"

type ExerciseValidatorInterface interface {
	ValidateExercise(exercise *dtos.CreateExerciseInput) error
}
