package entities

type TrainingActive struct {
	ID          string
	UserID      string
	TrainingID  string
	Repetitions int
	Weight      float64
	Time        float64
}

func NewTrainingActive(tu TrainingActive) (*TrainingActive, error) {
	trainingActive := &TrainingActive{
		ID:          tu.ID,
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
