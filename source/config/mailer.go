package config

import (
	"os"
	"sarana-dafa-ai-service/storage/env"
	"strconv"

	"gopkg.in/gomail.v2"
)

func GomailDialer() *gomail.Dialer {
	intEmailPort, _ := strconv.Atoi(os.Getenv(env.EMAIL_SMTP_PORT))
	return gomail.NewDialer(
		os.Getenv(env.EMAIL_SMTP_HOST),
		intEmailPort,
		os.Getenv(env.EMAIL_AUTH_EMAIL),
		os.Getenv(env.EMAIL_AUTH_PASSWORD),
	)
}
