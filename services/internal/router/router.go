package router

import (
	"database/sql"
	"services/internal/handler"
    "services/internal/middleware"
    "services/internal/service"
    "services/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func New(db *sql.DB) *fiber.App {
	app := fiber.New()

	userRepo := storage.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)


	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	protected := api.Group("/", middleware.AuthMiddleware)
	protected.Get("/profile", func(c *fiber.Ctx) error{
		user_id := c.Locals("user_id").(int)
		user_role := c.Locals("user_role").(string)

		return c.JSON(fiber.Map{
			"user_id": userID,    
			"role": userRole,
		})
	})

	return app
}
