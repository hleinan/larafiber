package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hleinan/larafiber/app/controllers/common"
	"github.com/hleinan/larafiber/app/services/helpers"
	"github.com/hleinan/larafiber/app/services/inertia"
)

func Index(c *fiber.Ctx) {
	dict := make(map[string]interface{})
	inertia.Render("Frontpage", c, dict)
}

func Me(c *fiber.Ctx) {
	user := helpers.CurrentUser(c)
	dict := common.Me(c, user.ID)
	inertia.Render("Me", c, dict)
}

func Verified(c *fiber.Ctx) {
	dict := make(map[string]interface{})
	inertia.Render("Verified", c, dict)
}
