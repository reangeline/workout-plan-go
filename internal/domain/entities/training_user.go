package entities

import "github.com/reangeline/workout-plan-go/pkg/entities"

type TrainingActive struct {
	UserID      entities.ID
	TrainingID  entities.ID
	Repetitions int
	Weight      float64
	Time        float64
}

func NewTrainingActive(tu TrainingActive) (*TrainingActive, error) {
	trainingActive := &TrainingActive{
		UserID:      tu.UserID,
		TrainingID:  tu.TrainingID,
		Repetitions: tu.Repetitions,
		Weight:      tu.Weight,
		Time:        tu.Time,
	}

	err := trainingActive.IsValid()

	if err != nil {
		return nil, err
	}

	return trainingActive, nil
}

func (e *TrainingActive) IsValid() error {

	return nil
}
