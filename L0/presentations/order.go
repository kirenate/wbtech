package presentations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"main.go/services"
)

type Presentation struct {
	service *services.Service
}

func NewPresentation(service *services.Service) *Presentation {
	return &Presentation{service: service}
}

func (r *Presentation) getOrder(c *fiber.Ctx) error {
	orderUID := c.Params("order_uid")

	order, err := r.service.GetOrder(c.UserContext(), orderUID)
	if err != nil {
		return errors.Wrap(err, "service failed to get order")
	}

	return c.JSON(order)
}
