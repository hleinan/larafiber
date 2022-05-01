package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ApiReq() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//sessions := store.GetSessions()
		//store := sessions.Get(c)

		//if store.Get("logged_in") == true {
		c.Next()
		//return
		//} else {
		//	c.Redirect("/login?next="+c.OriginalURL())
		//}
		return nil
	}
}
