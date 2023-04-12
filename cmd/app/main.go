package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/reangeline/workout-plan-go/configs"
	"github.com/reangeline/workout-plan-go/internal/infra/http/routes"

	_ "github.com/lib/pq"
	_ "github.com/reangeline/workout-plan-go/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Workout Plan API
// @version         1.0
// @description     API
// @termsOfService  http://swagger.io/terms/

// @contact.name   Renato Saraiva Angeline
// @contact.email  reangeline@hotmail.com

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(configs)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	fmt.Println("Postgres connected")

	if err != nil {
		panic(err)
	}

	uc, err := InitializeUserController(db)
	if err != nil {
		log.Fatalf("failed to initialize user controller: %v", err)
	}

	ec, err := InitializeExerciseController(db)
	if err != nil {
		log.Fatalf("failed to initialize exercise controller: %v", err)
	}

	tc, err := InitializeTrainingController(db)
	if err != nil {
		log.Fatalf("failed to initialize exercise controller: %v", err)
	}

	r := routes.InitializeUserRoutes(uc)
	r = routes.InitializeExerciseRoutes(ec)
	r = routes.InitializeTrainingRoutes(tc)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	err = http.ListenAndServe(":"+configs.WebServerPort, r)
	fmt.Println(err)

}
