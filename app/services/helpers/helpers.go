package helpers

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/hleinan/larafiber/app"
	"github.com/hleinan/larafiber/app/services/store"
)

func CurrentUser(c *fiber.Ctx) models.User {
	sessions := store.GetSessions()
	store := sessions.Get(c)
	id := store.Get("id")
	return models.GetUser(id)
}
