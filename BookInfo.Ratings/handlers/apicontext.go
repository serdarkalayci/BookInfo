package handlers

import (
	"bookinfo/ratings/dto"
	"bookinfo/ratings/logger"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
)

// KeyRating is a key used for the Rating object in the context
type KeyRating struct{}

// APIContext is the struct that has a logger and validation instance. It's the base for all handler functions
type APIContext struct {
	v *dto.Validation
}

// DBContext is the struct that has a MongoDB connection together with standard APIContext. It's used for handler functions which will use database
type DBContext struct {
	RedisClient  redis.Client
	DatabaseName int
	APIContext
}

// NewAPIContext returns a new APIContext handler with the given logger
func NewAPIContext(v *dto.Validation) *APIContext {
	return &APIContext{v}
}

// NewDBContext returns a new DBContext handler with the given logger
func NewDBContext(v *dto.Validation) *DBContext {
	// We try to get redisAddress value from the environment variables, if not found it falls back to local database
	redisAddress := os.Getenv("redisAddress")
	if redisAddress == "" {
		redisAddress = "mongodb://localhost:27017"
		logger.Log("redisAddress from Env not found, falling back to local DB", logger.DebugLevel)
	} else {
		logger.Log(fmt.Sprintf("redisAddress from Env is used: '%s'", redisAddress), logger.DebugLevel)
	}
	envDbName := os.Getenv("DatabaseName")
	databaseName := 0
	if envDbName == "" {
		logger.Log("DatabaseName from Env not found, falling back to default", logger.DebugLevel)
	} else {
		databaseName, err := strconv.Atoi(envDbName)
		if err != nil {
			databaseName = 0
			logger.Log("DatabaseName from Env is not in expected format, falling back to default", logger.DebugLevel)
		}
		logger.Log(fmt.Sprintf("DatabaseName from Env is used: '%s'", databaseName), logger.DebugLevel)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		logger.Log("An error occured while connecting to tha database", logger.ErrorLevel, err)
		log.Fatal("Cannot connect to database")
	}
	logger.Log("Connected to MongoDB!", logger.DebugLevel)
	return &DBContext{*client, databaseName, APIContext{v}}
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
