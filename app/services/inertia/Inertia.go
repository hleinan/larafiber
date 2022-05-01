package inertia

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	models "github.com/hleinan/larafiber/app"
	"github.com/hleinan/larafiber/app/services/store"
	"html"
	"html/template"
)

type Data struct {
	Component string                 `json:"component"`
	Props     map[string]interface{} `json:"props"`
	Url       string                 `json:"url"`
	Version   string                 `json:"version"`
}

func Render(component string, c *fiber.Ctx, props map[string]interface{}) {
	newProps := props
	newProps["auth"] = currentUser(c)
	newProps["error"] = errors(c)
	newProps["title"] = component

	deleteErrors(c)
	data := Data{
		Component: component,
		Props:     newProps,
		Url:       c.OriginalURL(),
		Version:   getLatestVersion(),
	}

	jsonOutput, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	if string(c.Request().Header.Peek("X-inertia")) == "true" {
		if string(c.Request().Header.Peek("X-Inertia-Version")) != data.Version {
			c.Set("X-Inertia-Location", data.Url)
			c.Status(409)
		} else {
			c.Set("X-inertia", "true")
			c.JSON(data)
		}

	} else {
		renderString := `<div id="app" data-page='` + html.EscapeString(string(jsonOutput)) + `'></div>`
		_ = c.Render("index", fiber.Map{
			"inertia": template.HTML(renderString),
			"version": data.Version,
			"title":   component,
		})
	}
}

func currentUser(c *fiber.Ctx) models.User {
	sessions := store.GetSessions()
	store := sessions.Get(c)
	id := store.Get("id")
	return models.GetUser(id)
}

func errors(c *fiber.Ctx) interface{} {
	sessions := store.GetSessions()
	store := sessions.Get(c)
	return store.Get("error")
}

func deleteErrors(c *fiber.Ctx) {
	sessions := store.GetSessions()
	store := sessions.Get(c)
	store.Set("error", "")
	defer store.Save()
}

func getLatestVersion() string {
	db := models.GetDB()
	var version models.Version

	if result := db.
		Last(&version); result.Error != nil {
		return "1"
	} else {
		return version.UUID
	}
}
