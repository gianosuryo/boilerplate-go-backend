package config

import (
	"fmt"
	"os"
	"sarana-dafa-ai-service/storage/env"
	"time"

	"github.com/gofiber/storage"
	"github.com/gofiber/storage/postgres/v3"
)

func NewStorage() storage.Storage {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv(env.DB_HOST),
		os.Getenv(env.DB_USER),
		os.Getenv(env.DB_PASSWORD),
		os.Getenv(env.DB_NAME),
		os.Getenv(env.DB_PORT),
	)
	storage := postgres.New(postgres.Config{
		ConnectionURI: dsn,
		Table:         "fiber_storage",
		Reset:         false,
		GCInterval:    10 * time.Second,
	})

	return storage
}
