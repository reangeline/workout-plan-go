package dtos

import "github.com/reangeline/workout-plan-go/pkg/entities"

type CreateTrainingInput struct {
	IDUser      entities.ID
	TrainigName string
}
