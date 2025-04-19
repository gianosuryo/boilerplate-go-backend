package main

import (
	"flag"
	"fmt"
	"os"
	"sarana-dafa-ai-service/config"
	"sarana-dafa-ai-service/controller"
	"sarana-dafa-ai-service/service"
	"sarana-dafa-ai-service/storage"
	"sarana-dafa-ai-service/storage/env"

	_ "sarana-dafa-ai-service/docs"

	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func main() {
	// Load data from os/file/string to var
	storage.InitStorage()
	config.InitCasbin()

	db := config.NewDatabase()
	val := config.NewValidator()

	// Additional CLI Command
	if cliMigrate(db) {
		return
	}

	// Fiber Definition
	app := config.NewFiber()
	config.SetRecover(app)
	config.SetCORS(app)
	config.SetAccessLogger(app)

	bumameAuthService := service.NewBumameAuthService(db)
	bumameB2BProductService := service.NewBumameB2BProductService(db)

	// Controller
	bumameAuthController := controller.NewBumameAuthController(bumameAuthService, val)
	bumameB2BProductController := controller.NewBumameB2BProductController(bumameB2BProductService, val)

	// Router
	config.BumameAuthRouter(app, bumameAuthController)
	config.BumameB2BProductRouter(app, bumameB2BProductController)

	// Only enable Swagger in non-production environments
	if os.Getenv(env.APP_ENV) == "dev" {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	app.Listen(":" + os.Getenv(env.APP_PORT))
}

// CLI Migrate
// Usage: go run server.go -migrate
func cliMigrate(db *gorm.DB) bool {
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	if *migrate {
		fmt.Println("Running database migrations...")
		config.MigrateTable(db)
		fmt.Println("Database migrations completed!")
		return true
	}
	return false
}
