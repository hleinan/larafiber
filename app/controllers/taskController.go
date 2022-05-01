package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	models "github.com/hleinan/larafiber/app"
	"github.com/hleinan/larafiber/app/services/helpers"
	"github.com/hleinan/larafiber/app/services/inertia"
	"log"
)

func Tasks(c *fiber.Ctx) {
	dict := make(map[string]interface{})

	db := models.GetDB()
	currentUser := helpers.CurrentUser(c)

	var tasks []models.Task
	db.
		Where("user_id = ?", currentUser.ID).
		Order("created_at desc").
		Find(&tasks)

	dict["tasks"] = tasks
	inertia.Render("Task/Index", c, dict)
}

func CreateTasks(c *fiber.Ctx) {
	task := new(models.Task)
	db := models.GetDB()

	if err := c.BodyParser(task); err != nil {
		log.Println(err)
	} else {
		currentUser := helpers.CurrentUser(c)
		task.UserID = currentUser.ID
		fmt.Println(task.Text, " ... ", task.UserID)
		db.Save(task)
	}

	c.Redirect("/task")
}

func ChangeTaskStatus(c *fiber.Ctx) {
	db := models.GetDB()
	currentUser := helpers.CurrentUser(c)

	var task models.Task
	db.
		Where("user_id", currentUser.ID).
		Where("uuid", c.Params("uuid")).
		First(&task)

	db.Model(&models.Task{}).Where("id = ?", task.ID).Update("done", !task.Done)

	c.Redirect("/task")
}
