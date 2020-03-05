package handlers

import (
	"bookinfo/ratings/dto"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyRating is a key used for the Rating object in the context
type KeyRating struct{}

// APIContext handler for getting and updating Ratings
type APIContext struct {
	l *log.Logger
	v *dto.Validation
}

// NewAPIContext returns a new APIContext handler with the given logger
func NewAPIContext(l *log.Logger, v *dto.Validation) *APIContext {
	return &APIContext{l, v}
}

// ErrInvalidRatingPath is an error message when the Rating path is not valid
var ErrInvalidRatingPath = fmt.Errorf("Invalid Path, path should be /Ratings/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getRatingID returns the Rating ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getRatingID(r *http.Request) int {
	// parse the Rating id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
