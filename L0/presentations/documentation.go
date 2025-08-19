package presentations

import (
	_ "embed"

	"github.com/gofiber/fiber/v2"
)

//go:embed docs/openapi.yaml
var openapi []byte

//go:embed docs/swagger.html
var swagger []byte

func (r *Presentation) openapi(c *fiber.Ctx) error {
	return c.Send(openapi)
}

func (r *Presentation) swagger(c *fiber.Ctx) error {
	return c.Type("html").Send(swagger)
}
