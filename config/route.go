package config

import (
	handler "docker-go/app/controllers"
	user_repository "docker-go/app/repositories/user"
	user_service "docker-go/app/services/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = SetupDatabase()

	// User
	userRepository user_repository.UserRepository = user_repository.NewUserRepository(db)
	UserService    user_service.UserService       = user_service.NewUserService(userRepository)
	userHandler    handler.UserHandler            = handler.NewUserHandler(UserService)
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/user")
	user.Get("/", userHandler.FindAll)
	user.Post("/", userHandler.Create)
	user.Get("/:id", userHandler.Find)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})
}
