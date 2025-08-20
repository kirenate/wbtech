package presentations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"main.go/repositories"
	"main.go/services"
)

type Presentation struct {
	service  *services.Service
	producer *services.Producer
}

func NewPresentation(service *services.Service, producer *services.Producer) *Presentation {
	return &Presentation{service: service, producer: producer}
}

func (r *Presentation) getOrder(c *fiber.Ctx) error {
	orderUID := c.Params("order_uid")

	order, err := r.service.GetOrder(c.UserContext(), orderUID)
	if err != nil {
		if errors.Is(err, repositories.ErrOrderDoesNotExist) {
			return &fiber.Error{
				Code:    fiber.StatusBadRequest,
				Message: "incorrect order",
			}
		}
		return errors.Wrap(err, "service failed to get order")
	}

	return c.JSON(order)
}

func (r *Presentation) generateOrders(c *fiber.Ctx) error {
	err := r.producer.SendMsg(c.UserContext())

	if err != nil {
		return errors.Wrap(err, "failed to send msg")
	}

	return c.JSON(&fiber.Map{"status": "success"})
}
