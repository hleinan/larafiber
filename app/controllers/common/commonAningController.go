package common

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/hleinan/larafiber/app"
)

func Me(c *fiber.Ctx, userID int64) map[string]interface{} {
	dict := make(map[string]interface{})
	db := models.GetDB()

	var complete_user models.User

	db.
		Where("id = ?", userID).
		First(&complete_user)

	var account models.Account

	db.
		Where("user_id = ?", userID).
		First(&account)

	dict["user"] = complete_user
	dict["account"] = account

	return dict
}
