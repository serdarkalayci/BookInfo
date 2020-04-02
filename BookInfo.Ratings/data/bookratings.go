package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Rating defines the structure for an API Rating
// swagger:model
type Rating struct {
	// the id of the book
	//
	// required: false
	// min: 1
	BookID int `json:"bookId" bson:"bookId"` // Unique identifier for the book

	// the rating of the book
	//
	// required: true
	// min: 0.01
	CurrentRating float32 `json:"currentRating" bson:"currentRating" validate:"required,gte=0"`

	// the rating of the book
	//
	// required: true
	// min: 0.01
	VoteCount int32 `json:"voteCount" bson:"voteCount" validate:"required,gte=0"`
}

// GetRatingByID returns a single Rating which matches the id from the
// database.
// If a Rating is not found this function returns a RatingNotFound error
func GetRatingByID(id int, dbClient mongo.Client, dbName string) (*Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := dbClient.Database(dbName).Collection("ratings")
	var rating Rating
	err := collection.FindOne(ctx, bson.M{"bookId": id}).Decode(&rating)
	if err != nil {
		return nil, err
	}
	return &rating, nil
}
