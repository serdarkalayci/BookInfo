package handlers

import (
	"bookinfo/ratings/dto"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// KeyRating is a key used for the Rating object in the context
type KeyRating struct{}

// APIContext is the struct that has a logger and validation instance. It's the base for all handler functions
type APIContext struct {
	l *log.Logger
	v *dto.Validation
}

// DBContext is the struct that has a MongoDB connection together with standard APIContext. It's used for handler functions which will use database
type DBContext struct {
	MongoClient  mongo.Client
	DatabaseName string
	APIContext
}

// NewAPIContext returns a new APIContext handler with the given logger
func NewAPIContext(l *log.Logger, v *dto.Validation) *APIContext {
	return &APIContext{l, v}
}

// NewDBContext returns a new DBContext handler with the given logger
func NewDBContext(l *log.Logger, v *dto.Validation) *DBContext {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// We try to get connectionstring value from the environment variables, if not found it falls back to local database
	connectionString := os.Getenv("ConnectionString")
	if connectionString == "" {
		connectionString = "mongodb://root:example@localhost:27017"
		l.Printf("ConnectionString from Env not found, falling back to local DB")
	} else {
		l.Printf("ConnectionString from Env is used: '%s'", connectionString)
	}
	databaseName := os.Getenv("DatabaseName")
	if databaseName == "" {
		databaseName = "cookiemonster"
		l.Printf("DatabaseName from Env not found, falling back to default")
	} else {
		l.Printf("DatabaseName from Env is used: '%s'", databaseName)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	l.Printf("Connected to MongoDB!")
	return &DBContext{*client, databaseName, APIContext{l, v}}
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
func getBookID(r *http.Request) int {
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
