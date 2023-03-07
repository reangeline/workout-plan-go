package entities_test

import (
	"testing"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyFields_WhenCreateExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := entities.Exercise{
		Name:        "",
		MuscleGroup: []string{},
		Description: "",
		UriGif:      "",
	}
	err := exercise.IsValid()
	assert.Error(t, err)
}
func TestGivenAnEmptyName_WhenCreateExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := entities.Exercise{
		MuscleGroup: []string{"peito"},
		Description: "des",
		UriGif:      "uri",
	}
	err := exercise.IsValid()
	assert.Error(t, err, "name is required")
}

func TestGivenAnEmptyMuscleGroup_WhenCreateExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := entities.Exercise{
		Name:        "supino",
		Description: "des",
		UriGif:      "uri",
	}
	err := exercise.IsValid()
	assert.Error(t, err, "muscle group is required")
}

func TestGivenAnEmptyDescription_WhenCreateExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := entities.Exercise{
		Name:        "supino",
		MuscleGroup: []string{"peito"},
		UriGif:      "uri",
	}
	err := exercise.IsValid()
	assert.Error(t, err, "descritpion is required")
}

func TestGivenAnEmptyUriGif_WhenCreateExercise_ThenShouldReceiveAnError(t *testing.T) {
	exercise := entities.Exercise{
		Name:        "supino",
		MuscleGroup: []string{"peito"},
		Description: "des",
	}
	err := exercise.IsValid()
	assert.Error(t, err, "uri gif is required")
}

func TestGivenAValidParams_WhenICallNewExercise_ThenIShouldReceiveCreateExerciseWithAllParams(t *testing.T) {
	exercise := entities.Exercise{
		Name:        "supino",
		MuscleGroup: []string{"peito"},
		Description: "des",
		UriGif:      "uri",
	}
	assert.Equal(t, "supino", exercise.Name)
	assert.Equal(t, []string{"peito"}, exercise.MuscleGroup)
	assert.Equal(t, "des", exercise.Description)
	assert.Equal(t, "uri", exercise.UriGif)
	assert.Nil(t, exercise.IsValid())

}

func TestGivenAValidParams_WhenICallNewExerciseFunc_ThenIShouldReceiveCreateExerciseWithAllParams(t *testing.T) {
	exer := entities.Exercise{
		Name:        "supino",
		MuscleGroup: []string{"peito"},
		Description: "des",
		UriGif:      "uri",
	}

	exercise, err := entities.NewExercise(exer)
	assert.Nil(t, err)

	assert.Equal(t, "supino", exercise.Name)
	assert.Equal(t, []string{"peito"}, exercise.MuscleGroup)
	assert.Equal(t, "des", exercise.Description)
	assert.Equal(t, "uri", exercise.UriGif)

}
