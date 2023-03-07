package entities

import "errors"

type Training struct {
	ID          string
	UserID      string
	TrainigName string
	Exercises   []*Exercise
}

func NewTraining(t Training) (*Training, error) {
	training := &Training{
		ID:          t.ID,
		UserID:      t.UserID,
		TrainigName: t.TrainigName,
		Exercises:   t.Exercises,
	}

	err := t.IsValid()

	if err != nil {
		return nil, err
	}

	return training, nil
}

func (t *Training) IsValid() error {
	if t.ID == "" {
		return errors.New("id is required")
	}

	if t.UserID == "" {
		return errors.New("user id is required")
	}

	if t.TrainigName == "" {
		return errors.New("training name is required")
	}
	return nil
}
