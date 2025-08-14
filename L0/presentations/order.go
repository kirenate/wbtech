package presentations

import (
	"github.com/gofiber/fiber/v2"
	"main.go/services"
)

type Presentation struct {
	service *services.Service
}

func NewPresentation(service *services.Service) *Presentation {
	return &Presentation{service: service}
}

func (r *Presentation) getOrders(c *fiber.Ctx) error {
	orderUID := c.Params("order_uid")

	return c.JSON(fiber.Map{"order_uid": orderUID})
}
