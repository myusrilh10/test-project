package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/myusrilh10/test-project/rest-api/database"
	"github.com/myusrilh10/test-project/rest-api/handlers"
)

func main() {
	// init app

	err := initApp()
	if err != nil {
		panic(err)
	}

	defer database.CloseMongoDB()

	app := generateApp()

	// get the port from env
	port := os.Getenv("PORT")

	app.Listen(":" + port)
}

func generateApp() *fiber.App {
	app := fiber.New()

	// create health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	// create the library group and routes
	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.GetLibraries)
	libGroup.Post("/", handlers.CreateLibrary)

	return app
}

func initApp() error {
	// setup env
	err := loadEnv()
	if err != nil {
		return err
	}
	// setup database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	return nil
}

func loadEnv() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}
