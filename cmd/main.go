package main

import (
	"fmt"

	"github.com/lunarr-app/lunarr-go/internal/config"
	"github.com/lunarr-app/lunarr-go/internal/handlers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func main() {
	// Get the config
	cfg := config.Get()

	// Create a new Iris application
	app := iris.New()

	// Set up a custom logger
	customLogger := logger.New(logger.Config{
		Status:           true,
		IP:               true,
		Method:           true,
		Path:             true,
		PathAfterHandler: true,
		Query:            true,
	})

	// Register the custom logger as middleware
	app.Use(customLogger)

	// Create a sub-router for all API routes
	api := app.Party("/api")

	// Register api routes
	api.Get("/", handlers.RootHandler)
	api.Post("/signup", handlers.SignupHandler)

	// Start the server on the specified port
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	app.Run(iris.Addr(addr))
}
