package config

import (
	"sarana-dafa-ai-service/controller"

	"github.com/gofiber/fiber/v2"
)

func BumameAuthRouter(app *fiber.App, cont controller.BumameAuthController) {
	app.Post("/auth/login", cont.Login)
	app.Get("/auth/read-token", cont.ReadToken)
}

func BumameB2BProductRouter(app *fiber.App, cont controller.BumameB2BProductController) {
	app.Put("/b2b-product/bulk-update", cont.BulkUpdate)

	app.Get("/b2b-product", cont.FindAll)
	app.Get("/b2b-product/:b2b_product_id", cont.FindById)
	app.Post("/b2b-product", cont.Create)
	app.Put("/b2b-product/:b2b_product_id", cont.Update)
	app.Delete("/b2b-product/:b2b_product_id", cont.Delete)

	app.Post("/b2b-product/generate-slugs", cont.GenerateSlugs)
}
