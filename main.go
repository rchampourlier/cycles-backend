//go:generate goagen bootstrap -d github.com/rchampourlier/cycles-backend/design

package main

import (
	"log"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/rchampourlier/cycles-backend/app"
	"github.com/rchampourlier/cycles-backend/commands"
	"github.com/rchampourlier/cycles-backend/queries"
)

func main() {
	// Configuration
	bucket := os.Getenv("AWS_BUCKET")
	if len(bucket) == 0 {
		log.Fatalf("AWS_BUCKET environment variable must be provided")
	}

	// Create commands and queries
	stateCommands := commands.NewStateCommands(bucket)
	stateQueries := queries.NewStateQueries(bucket)

	// Create service
	service := goa.New("cycles")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "state" controller
	c := NewStateController(service, stateCommands, stateQueries)
	app.MountStateController(service, c)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}
}
