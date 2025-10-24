package main

import (
	"log"
	"user-management-api/config"
	"user-management-api/internal/handler"
	"user-management-api/internal/logger"
	"user-management-api/internal/repository"
	"user-management-api/internal/routes"
	"user-management-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func main() {
	if err := logger.InitLogger(); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}

	cfg := config.Load()

	db, err := repository.NewDatabase(cfg)
	if err != nil {
		logger.Log.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Log.Error("Unhandled error", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		},
	})

	app.Use(recover.New())

	routes.SetupRoutes(app, userHandler)

	logger.Log.Info("Starting server on port " + cfg.ServerPort)
	if err := app.Listen(cfg.ServerPort); err != nil {
		logger.Log.Fatal("Failed to start server", zap.Error(err))
	}
}
