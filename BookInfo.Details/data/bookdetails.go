package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ErrDetailNotFound is an error raised when a detail can not be found in the database
var ErrDetailNotFound = fmt.Errorf("Detail not found")

// Detail defines the structure for an Book detail
// swagger:model
type Detail struct {
	// the id of the book
	//
	// required: false
	// min: 1
	BookID int `json:"bookId" bson:"bookId"` // Unique identifier for the book

	// the name of the book
	//
	// required: true
	Name string `json:"name" bson:"name" validate:"required"`

	// the ISBN of the book
	//
	// required: true
	ISBN string `json:"isbn" bson:"isbn" validate:"required"`

	// the author of the book
	//
	// required: true
	Author string `json:"author" bson:"author" validate:"required"`

	// the publish date of the book
	//
	// required: true
	PublishDate time.Time `json:"publishDate" bson:"publishDate" validate:"required"`

	// the price of the book
	//
	// required: true
	// min: 0.01
	Price float64 `json:"price" bson:"price" validate:"required,gte=0"`

	// the number of books in the stock
	//
	// required: false
	// min: 0
	CurrentStock int `json:"currentStock"` // the number of books in the stock
}

// GetDetailByID returns a single Detail which matches the id from the
// database.
// If a Detail is not found this function returns a DetailNotFound error
func GetDetailByID(id int, dbClient mongo.Client, dbName string) (*Detail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := dbClient.Database(dbName).Collection("details")
	var detail Detail
	err := collection.FindOne(ctx, bson.M{"bookId": id}).Decode(&detail)
	if err != nil {
		return nil, err
	}
	return &detail, nil
}



