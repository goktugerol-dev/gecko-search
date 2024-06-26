package routes

import (
	"gecko-crawl/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return render(c, views.Home())
	})
}

// func SetRoutes(app *fiber.App) {
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.JSON(fiber.Map{
// 			"Message": "Helo, World!",
// 		})
// 	})
// }
