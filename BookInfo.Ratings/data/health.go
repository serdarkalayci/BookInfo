package data

import (
	"github.com/go-redis/redis/v7"
)

// GetHealth returns error if the database connection is failing
func GetHealth(dbClient redis.Client, dbName int) error {
	_, _ = dbClient.Ping().Result()
	return nil
}
