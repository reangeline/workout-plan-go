package entities

import (
	"errors"

	"github.com/reangeline/workout-plan-go/pkg/entities"
)

type Exercise struct {
	IDExercise   entities.ID
	NameExercise string
	MuscleGroup  []string
	Description  string
	UriGif       string
}

func NewExercise(name_exercise string, muscle_group []string, description string, uri_gif string) (*Exercise, error) {
	exercise := &Exercise{
		NameExercise: name_exercise,
		MuscleGroup:  muscle_group,
		Description:  description,
		UriGif:       uri_gif,
	}

	err := exercise.IsValid()

	if err != nil {
		return nil, err
	}

	exercise.AddId()

	return exercise, nil
}

func (u *Exercise) AddId() {
	u.IDExercise = entities.NewID()
}

func (e *Exercise) IsValid() error {
	if e.NameExercise == "" {
		return errors.New("name is required")
	}

	if len(e.MuscleGroup) == 0 {
		return errors.New("muscle group is required")

	}

	if e.Description == "" {
		return errors.New("descritpion is required")
	}

	if e.UriGif == "" {
		return errors.New("uri gif is required")
	}
	return nil
}
