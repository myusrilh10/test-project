package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/myusrilh10/test-project/rest-api/database"
	"github.com/myusrilh10/test-project/rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

type newBookDTO struct {
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	ISBN      string `json:"isbn" bson:"isbn"`
	LibraryId string `json:"libraryId" bson:"libraryId"`
}

func CreateBook(c *fiber.Ctx) error {
	createData := new(newBookDTO)

	if err := c.BodyParser(createData); err != nil {
		return err
	}

	bookCollection := database.GetCollection("books")
	_, err := bookCollection.InsertOne(context.TODO(), createData)
	if err != nil {
		return err
	}

	// get the collection reference
	coll := database.GetCollection("libraries")

	// get the filter
	filter := bson.M{"id": createData.LibraryId}
	nBookData := models.Book{
		Title:  createData.Title,
		Author: createData.Author,
		ISBN:   createData.ISBN,
	}
	updatePayload := bson.M{"$push": bson.M{"books": nBookData}}

	//update the library
	_, salah := coll.UpdateOne(context.TODO(), filter, updatePayload)
	if salah != nil {
		return salah
	}

	return c.SendString("book created successfully")
}
