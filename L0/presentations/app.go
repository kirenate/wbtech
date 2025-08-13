package presentations

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

func (r *Presentation) BuildApp() *fiber.App {
	app := fiber.New(fiber.Config{})
	app.Use(recover2.New(recover2.Config{EnableStackTrace: true}))
	return app
}
