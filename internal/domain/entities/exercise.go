package entities

import (
	"errors"
)

type Exercise struct {
	Name        string
	MuscleGroup []string
	Description string
	UriGif      string
}

func NewExercise(name string, muscle_group []string, description string, uri_gif string) (*Exercise, error) {
	exercise := &Exercise{
		Name:        name,
		MuscleGroup: muscle_group,
		Description: description,
		UriGif:      uri_gif,
	}

	err := exercise.IsValid()

	if err != nil {
		return nil, err
	}

	return exercise, nil
}

func (e *Exercise) IsValid() error {
	if e.Name == "" {
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
