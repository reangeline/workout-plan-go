package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/reangeline/workout-plan-go/configs"
	"github.com/reangeline/workout-plan-go/internal/infra/http/routes"

	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Meu API REST
// @description Esta é uma API REST de exemplo
// @version 1
// @host localhost:8000
// @BasePath /
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(configs)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	fmt.Println("Postgres ok")

	if err != nil {
		panic(err)
	}

	uc, err := InitializeUserController(db)
	if err != nil {
		log.Fatalf("failed to initialize user controller: %v", err)
	}

	r := routes.InitializeUserRoutes(uc)
	r.Get("/docs/*", httpSwagger.WrapHandler)

	err = http.ListenAndServe(":"+configs.WebServerPort, r)
	fmt.Println(err)

}
