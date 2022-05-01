package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hleinan/larafiber/app/services/inertia"
	"strings"
)

func Dashboard(c *fiber.Ctx) {
	dict := make(map[string]interface{})
	uuidGen := uuid.Must(uuid.NewRandom()).String()
	dict["uuid"] = strings.ReplaceAll(uuidGen, "-", "")
	dict["host"] = string(c.Context().URI().FullURI())
	dict["context"] = c.GetReqHeaders()
	dict["context_2"] = c.Get("Host")
	dict["domain"] = c.BaseURL()
	dict["domain2"] = c.Request().Host()
	dict["domain_url"] = c.Request().URI().PathOriginal()
	inertia.Render("Admin/Dashboard", c, dict)
}
