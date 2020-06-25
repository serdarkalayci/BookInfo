package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// GetHealth chects if it can connect to the database and returns error if it's not possible
func GetHealth(dbClient mongo.Client, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := dbClient.Ping(ctx, readpref.Primary())
	return err
}
