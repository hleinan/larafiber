package render

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Data struct {
	Props   map[string]interface{} `json:"props"`
	Url     string                 `json:"url"`
	Version string                 `json:"version"`
}

func Render(c *fiber.Ctx, props map[string]interface{}) error {
	data := Data{
		Props:   props,
		Url:     c.OriginalURL(),
		Version: "2",
	}
	if string(c.Request().Header.Peek("X-Version")) != data.Version {
		// c.Status(409) deal with this later
	}

	fmt.Println("Version: ", string(c.Request().Header.Peek("X-Version")))
	c.Type("json", "utf-8")
	return c.JSON(data)
}
