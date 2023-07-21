package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/amopromo/gochip/api"
	"github.com/amopromo/gochip/api/handler"
	"github.com/amopromo/gochip/repository/sql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open(
		"postgres",
		"postgres://postgres:postgres@localhost:5432/chip?sslmode=disable",
	)

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	destinationRepository := sql.NewDestinationRepository(db)
	destinationHandler := &handler.DestinationHandler{DestinationRepository: destinationRepository}

	apiclientRepository := sql.NewApiClientRepository(db)

	// init gin route here
	r := gin.Default()

	// setup middleware
	api.SetupMiddleware(r, apiclientRepository)

	// setup routes
	api.SetupRouter(r, destinationHandler)

	r.Run(":8002")
}
