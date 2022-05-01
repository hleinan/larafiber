package middleware

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/hleinan/larafiber/app"
	"github.com/hleinan/larafiber/app/services/store"
)

func AdminReq() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessions := store.GetSessions()
		store := sessions.Get(c)

		id := store.Get("id")
		user := models.GetUser(id)

		if user.Admin == true {
			c.Next()
			return nil
		} else {
			c.Redirect("/")
		}
		return nil
	}
}
