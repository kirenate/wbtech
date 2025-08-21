package presentations

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/teadove/teasutils/fiber_utils"
	"time"
)

func (r *Presentation) BuildApp() *fiber.App {
	app := fiber.New(fiber.Config{
		Immutable:    true,
		ErrorHandler: fiber_utils.ErrHandler()})

	app.Use(fiber_utils.MiddlewareLogger())
	app.Use(fiber_utils.MiddlewareCtxTimeout(29 * time.Second))
	app.Use(recover2.New(recover2.Config{EnableStackTrace: true}))
	
	app.Get("/order/:order_uid", r.getOrder)
	app.Get("/generate-orders", r.generateOrders)

	app.Get("/docs", r.swagger)
	app.Get("/openapi.yaml", r.openapi)
	return app
}
