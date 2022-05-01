package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hleinan/larafiber/app/services/store"
)

func AuthReq() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessions := store.GetSessions()
		store := sessions.Get(c)

		if store.Get("logged_in") == true {
			c.Next()
			return nil
		} else {
			nextString := ""
			if c.Method() == "GET" {
				nextString = "?next=" + c.OriginalURL()
			}
			c.Redirect("/login" + nextString)
		}
		return nil
	}
}
