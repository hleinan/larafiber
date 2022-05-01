package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func Timer() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		// start timer
		start := time.Now()
		// next routes
		c.Next()
		// stop timer
		stop := time.Now()
		// Do something with response
		c.Append("Server-Timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))
	}
}
