package data

import (
	"errors"
	"fmt"
	"reflect"
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
func GetRatingByID(id int, dbClient redis.Client, dbName int) (*Rating, error) {
	result, err := dbClient.HGetAll(strconv.Itoa(id)).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	}
	var rating Rating
	rating.FillStruct(result)
	return &rating, nil
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func (r *Rating) FillStruct(m map[string]string) error {
	for k, v := range m {
		err := SetField(r, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
