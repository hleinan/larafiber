package controllers

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/hleinan/larafiber/app"
	"github.com/hleinan/larafiber/app/services/inertia"
)

func Show(c *fiber.Ctx) {
	db := models.GetDB()
	var user models.User
	db.
		// Preload("Entities").
		Preload("UserQueues").
		Where("uuid", c.Params("uuid")).
		First(&user)
	dict := make(map[string]interface{})
	dict["user"] = user
	inertia.Render("User", c, dict)
}
