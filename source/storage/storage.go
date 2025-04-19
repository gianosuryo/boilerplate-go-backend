package storage

import (
	"os"
	"sarana-dafa-ai-service/storage/env"
	"strings"
)

var ReservedSubdomain []string

func InitStorage() {
	env.InitEnv()

	// Load file reserverd-subdomain
	data, _ := os.ReadFile("./reserved-subdomain.txt")
	ReservedSubdomain = strings.Split(string(data), "\n")
}
