package entities_test

import (
	"testing"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	pkg_entities "github.com/reangeline/workout-plan-go/pkg/entities"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyFields_WhenCreateExercise_ThenShouldReceiveAnError(t *testing.T) {

	exercise := entities.Exercise{
		NameExercise: "",
		MuscleGroup:  []string{},
		Description:  "",
		UriGif:       "",
	}

	err := exercise.IsValid()
	assert.Error(t, err)
}

func TestGivenAValidParams_WhenICallNewExercise_ThenIShouldReceiveCreateExerciseWithAllParams(t *testing.T) {

	exercise := entities.Exercise{
		NameExercise: "Supino Reto",
		MuscleGroup:  []string{"peito"},
		Description:  "Supino no banco com barra",
		UriGif:       "http://teste.gif",
	}

	assert.Equal(t, "Supino Reto", exercise.NameExercise)
	assert.Equal(t, "Supino no banco com barra", exercise.Description)
	assert.Equal(t, "http://teste.gif", exercise.UriGif)

	err := exercise.IsValid()
	assert.NoError(t, err)
}

func TestGivenAValidParams_WhenICallNewExerciseFunc_ThenIShouldReceiveCreateExerciseWithAllParams(t *testing.T) {

	u, err := entities.NewExercise("supino", []string{"peito"}, "supino no banco", "http://teste.gif")
	assert.Nil(t, err)

	uuid, _ := pkg_entities.ParseID("00000000-0000-0000-0000-000000000000")

	assert.NotNil(t, u.IDExercise)
	assert.NotEqual(t, u.IDExercise, uuid)
	assert.Equal(t, "supino", u.NameExercise)
	assert.Equal(t, []string{"peito"}, u.MuscleGroup)
	assert.Equal(t, "supino no banco", u.Description)
	assert.Equal(t, "http://teste.gif", u.UriGif)
}
