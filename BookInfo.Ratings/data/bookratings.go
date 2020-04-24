package data

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v7"
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
	CurrentRating float64 `json:"currentRating" bson:"currentRating" validate:"required,gte=0"`

	// the rating of the book
	//
	// required: true
	// min: 0.01
	VoteCount int `json:"voteCount" bson:"voteCount" validate:"required,gte=0"`
}

// GetRatingByID returns a single Rating which matches the id from the
// database.
// If a Rating is not found this function returns a RatingNotFound error
func GetRatingByID(id int, dbClient redis.Client, dbName int) (*Rating, error) {
	result, err := dbClient.HGetAll(fmt.Sprintf("bookId:%d", id)).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	}
	var rating Rating
	rating.BookID = id
	rating.CurrentRating, err = strconv.ParseFloat(result["currentRating"], 64)
	if err != nil {
		rating.CurrentRating = 0
	}
	rating.VoteCount, err = strconv.Atoi(result["voteCount"])
	if err != nil {
		rating.VoteCount = 0
	}
	return &rating, nil
}

