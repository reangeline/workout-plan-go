package protocols

import "github.com/reangeline/workout-plan-go/internal/dtos"

type TrainingValidatorInterface interface {
	ValidateTraining(exercise *dtos.CreateTrainingInput) error
}
