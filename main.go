package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/template/html"
	models "github.com/hleinan/larafiber/app"
	"github.com/hleinan/larafiber/app/services/database"
	"github.com/hleinan/larafiber/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := html.New(os.Getenv("template_path"), ".html")

	app := fiber.New(
		fiber.Config{
			Prefork:       false,
			CaseSensitive: true,
			StrictRouting: false,
			ServerHeader:  "Fiber",
			Views:         engine,
		})

	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	app.Static("/static", "./public")

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	_ = models.GetDB()

	file, err := os.OpenFile("./log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Output: file,
	}))
	
	routes.WebRoutes(app)

	// Running database backup script
	database.Backup()

	log.Fatal(app.Listen(":" + os.Getenv("port")))

	models.CloseDB()
}
