package main

import (
	"log"

	"go-users-api/internal/config"
	"go-users-api/internal/handler"
	"go-users-api/internal/logger"
	"go-users-api/internal/repository"
	"go-users-api/internal/routes"
	"go-users-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logg, err := logger.New()
	if err != nil {
		log.Fatal(err)
	}
	defer logg.Sync()

	// Initialize DB connection
	db, err := config.NewMySQLConnection()
	if err != nil {
		logg.Fatal("Failed to connect to DB", zap.Error(err))
	}
	defer db.Close()

	logg.Info("Database connected successfully")

	// Initialize Fiber app
	app := fiber.New()

	// ---- Dependency wiring (IMPORTANT) ----
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Register routes
	routes.RegisterUserRoutes(app, userHandler)

	// Start server
	logg.Info("Server started", zap.String("port", "8080"))
	log.Fatal(app.Listen(":8080"))
}
