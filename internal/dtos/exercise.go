package dtos

import "github.com/reangeline/workout-plan-go/pkg/entities"

type CreateExerciseInput struct {
	NameExercise string   `json:"name"`
	MuscleGroup  []string `json:"muscle_group"`
	Description  string   `json:"description"`
	UriGif       string   `json:"uri_gif"`
}

type CreateExerciseOutput struct {
	IDExercise   entities.ID `json:"id_exercise"`
	NameExercise string      `json:"name"`
	MuscleGroup  []string    `json:"muscle_group"`
	Description  string      `json:"description"`
	UriGif       string      `json:"uri_gif"`
}
