package data

import (
	"bookinfo/ratings/dto"
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

// Ratings defines a slice of Rating
type Ratings []*Rating

// GetRatings returns all Ratings from the database
func GetRatings() Ratings {
	return RatingList
}

// GetRatingByID returns a single Rating which matches the id from the
// database.
// If a Rating is not found this function returns a RatingNotFound error
func GetRatingByID(id int, dbClient mongo.Client, dbName string) (*Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := dbClient.Database(dbName).Collection("Ratings")
	var rating Rating
	err := collection.FindOne(ctx, bson.M{"bookId": id}).Decode(&rating)
	if err != nil {
		return &rating, nil
	}
	return nil, err
}

// UpdateRating replaces a Rating in the database with the given
// item.
func UpdateRating(rat dto.NewRating) {
	i := findIndexByRatingID(rat.BookID)
	if i == -1 {
		// add new rating for the book
		newRating := &Rating{BookID: rat.BookID, CurrentRating: rat.Rating, VoteCount: 1}
		RatingList = append(RatingList, newRating)
	} else {
		// update the Rating in the DB
		newRating := ((RatingList[i].CurrentRating * (float32)(RatingList[i].VoteCount)) + rat.Rating) / (float32)(RatingList[i].VoteCount+1)
		RatingList[i].CurrentRating = newRating
		RatingList[i].VoteCount = RatingList[i].VoteCount + 1
	}
}

// findIndex finds the index of a Rating in the database
// returns -1 when no Rating can be found
func findIndexByRatingID(id int) int {
	for i, p := range RatingList {
		if p.BookID == id {
			return i
		}
	}

	return -1
}

var RatingList = []*Rating{
	&Rating{
		BookID:        1,
		CurrentRating: 2.45,
		VoteCount:     5,
	},
	&Rating{
		BookID:        2,
		CurrentRating: 1.99,
		VoteCount:     7,
	},
}
