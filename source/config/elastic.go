package config

import (
	"fmt"
	"log"
	"os"
	"sarana-dafa-ai-service/storage/env"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewElastic() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv(env.LOG_ELASTIC_HOST),
		},
		Username: os.Getenv(env.LOG_ELASTIC_USERNAME),
		Password: os.Getenv(env.LOG_ELASTIC_PASSWORD),
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	} else {
		fmt.Println("Elastic connected")
	}

	return es
}

func AddIndexElastic() {

}
