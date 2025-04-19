package config

import (
	"os"
	"sarana-dafa-ai-service/storage/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
)

func InitGoth(storage fiber.Storage) {
	goth.UseProviders(
		google.New(os.Getenv(env.GOOGLE_CLIENT_ID), os.Getenv(env.GOOGLE_SECRET), os.Getenv(env.GOOGLE_SSO_REDIRECT),
			"email", "profile"),
	)

	goth_fiber.SessionStore = session.New(session.Config{
		KeyLookup:      "cookie:dinosaurus",
		CookieHTTPOnly: true,
		Storage:        storage,
	})

}
