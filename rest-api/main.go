package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/myusrilh10/test-project/rest-api/database"
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
