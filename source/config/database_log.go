package config

import (
	"context"
	"log"
	"os"
	"sarana-dafa-ai-service/storage/env"
	"time"

	elastic "github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDatabaseLog() *mongo.Database {

	// Create connection string
	cs := "mongodb://" + os.Getenv(env.LOG_DB_USER) + ":" + os.Getenv(env.LOG_DB_PASSWORD) + "@" + os.Getenv(env.LOG_DB_HOST) + ":" + os.Getenv(env.LOG_DB_PORT) + "/?retryWrites=true&w=majority"

	// Create connection context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect mongodb database
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cs))
	if err != nil {
		panic(err)
	}

	// Return client database
	db := client.Database(os.Getenv(env.LOG_DB_NAME))
	return db
}

func NewDatabaseLogElastic() *elastic.Client {
	client, err := elastic.NewClient(elastic.Config{
		Addresses: []string{os.Getenv(env.LOG_ELASTIC_HOST)},
		Username:  os.Getenv(env.LOG_ELASTIC_USERNAME),
		Password:  os.Getenv(env.LOG_ELASTIC_PASSWORD),
	})

	if err != nil {
		log.Panic(err)
	}

	return client
}
