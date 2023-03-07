package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/reangeline/workout-plan-go/configs"
	"github.com/reangeline/workout-plan-go/internal/infra/http/routes"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(configs)
	}

	db, err := sql.Open("mysql", "user:password@tcp(host:port)/database")
	if err != nil {
		panic(err)
	}

	uc, err := InitializeUserController(db)
	if err != nil {
		log.Fatalf("failed to initialize user controller: %v", err)
	}

	r := routes.InitializeUserRoutes(uc)

	http.ListenAndServe(configs.WebServerPort, r)

}
