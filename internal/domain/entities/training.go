package entities

import (
	"errors"

	"github.com/reangeline/workout-plan-go/pkg/entities"
)

type Training struct {
	TrainingID  entities.ID
	UserID      entities.ID
	TrainigName string
}

func NewTraining(UserID entities.ID, trainig_name string) (*Training, error) {
	training := &Training{
		UserID:      UserID,
		TrainigName: trainig_name,
	}

	err := training.IsValid()

	if err != nil {
		return nil, err
	}

	return training, nil
}

func (u *Training) AddId() {
	u.TrainingID = entities.NewID()
}

func (t *Training) IsValid() error {

	if t.TrainigName == "" {
		return errors.New("training name is required")
	}
	return nil
}
