package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	jwtware "github.com/gofiber/jwt/v2"
	"os"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api/v1/")

	cacheApi := api
	cacheApi.Use(cache.New())

	// JWT Middleware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("jwt_token")),
	}))
}
