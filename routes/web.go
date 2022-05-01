package routes

import (
	"github.com/gofiber/fiber/v2"
	Controller "github.com/hleinan/larafiber/app/controllers"
	AdminController "github.com/hleinan/larafiber/app/controllers/admin"
	AuthController "github.com/hleinan/larafiber/app/controllers/authentification"
	Middeleware "github.com/hleinan/larafiber/app/middelware"
)

func WebRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		Controller.Index(c)
		return nil
	})

	app.Get("/verifyemail/:uuid", func(c *fiber.Ctx) error {
		AuthController.VerifyEmailExternal(c)
		return nil
	})

	app.Get("/verified", func(c *fiber.Ctx) error {
		Controller.Verified(c)
		return nil
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		AuthController.ShowLogin(c)
		return nil
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		AuthController.Login(c)
		return nil
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		AuthController.Logout(c)
		return nil
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		AuthController.ShowRegister(c)
		return nil
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		AuthController.Register(c)
		return nil
	})

	app.Get("/forgot", func(c *fiber.Ctx) error {
		AuthController.Forgot(c)
		return nil
	})

	app.Post("/forgot", func(c *fiber.Ctx) error {
		AuthController.ForgotStore(c)
		return nil
	})

	app.Get("/resetpassword/:uuid", func(c *fiber.Ctx) error {
		AuthController.ResetPassword(c)
		return nil
	})

	app.Post("/resetpassword/:uuid", func(c *fiber.Ctx) error {
		AuthController.ResetPasswordStore(c)
		return nil
	})

	/* User logged in */

	user := app.Group("/", Middeleware.AuthReq())

	app.Get("/task", func(c *fiber.Ctx) error {
		Controller.Tasks(c)
		return nil
	})

	app.Post("/task", func(c *fiber.Ctx) error {
		Controller.CreateTasks(c)
		return nil
	})

	app.Post("/task/change/:uuid", func(c *fiber.Ctx) error {
		Controller.ChangeTaskStatus(c)
		return nil
	})

	user.Get("/user/me", func(c *fiber.Ctx) error {
		Controller.Me(c)
		return nil
	})

	/* Admin */

	admin := app.Group("/admin/", Middeleware.AuthReq(), Middeleware.AdminReq())

	admin.Get("/", func(c *fiber.Ctx) error {
		AdminController.Dashboard(c)
		return nil
	})
}
