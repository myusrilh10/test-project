package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/myusrilh10/test-project/rest-api/database"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// init app

	err := initApp()
	if err != nil {
		panic(err)
	}

	defer database.CloseMongoDB()

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		// write a todo to the database
		sampleDoc := bson.M{"name": "sample todo"}
		collection := database.GetCollection("sambal")
		nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("error inserting todo")
		}
		return c.JSON(nDoc)
	})

	app.Listen(":3000")
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
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
