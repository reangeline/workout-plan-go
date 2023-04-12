package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/reangeline/workout-plan-go/internal/infra/http/controllers"
)

func InitializeTrainingRoutes(uc *controllers.TrainingController) chi.Router {
	r := chi.NewRouter()

	r.Route("/training", func(r chi.Router) {
		r.Post("/", uc.CreateTraining)
	})

	return r
}
