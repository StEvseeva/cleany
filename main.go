package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/StEvseeva/cleany/internal/db"
	"github.com/StEvseeva/cleany/internal/repository"
	"github.com/StEvseeva/cleany/internal/server"
	"github.com/StEvseeva/cleany/internal/service"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()

	swagger, err := server.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Initialize database
	dbConfig := db.ConfigFromEnv()
	database, err := db.NewPostgresDB(dbConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %s", err)
		os.Exit(1)
	}
	defer database.Close(context.Background())

	// Initialize repositories
	bookingRepo := repository.NewBookingRepository(database.GetDB())
	cleanerRepo := repository.NewCleanerRepository(database.GetDB())
	roomRepo := repository.NewRoomRepository(database.GetDB())
	cleaningOrderRepo := repository.NewCleaningOrderRepository(database.GetDB())

	// Initialize services
	bookingService := service.NewBookingService(bookingRepo, roomRepo)
	cleanerService := service.NewCleanerService(cleanerRepo)
	roomService := service.NewRoomService(roomRepo)
	cleaningOrderService := service.NewCleaningOrderService(cleaningOrderRepo, bookingRepo, cleanerRepo)

	// Create an instance of our handler which satisfies the generated interface
	api := server.NewServer(bookingService, cleanerService, roomService, cleaningOrderService)

	// This is how you set up a basic Echo router
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	// e.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	server.RegisterHandlers(e, api)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", *port)))
}
