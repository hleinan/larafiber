package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hleinan/larafiber/app/services/store"
)

func Guest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessions := store.GetSessions()
		store := sessions.Get(c)

		if store.Get("logged_in") == true {
			store.Regenerate()
		}

		return c.Next()
		return nil
	}
}
