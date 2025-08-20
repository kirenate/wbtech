package presentations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"main.go/repositories"
	"main.go/services"
	"main.go/utils"
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
	if errors.Is(err, repositories.ErrOrderDoesNotExist) {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "incorrect order",
		}
	}
	if err != nil {
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: errors.Wrap(err, "service failed to get order").Error(),
		}
	}

	return c.JSON(order)
}

func (r *Presentation) generateOrders(c *fiber.Ctx) error {
	producer := utils.NewProducer()
	err := producer.SendMsg(c.UserContext())

	if err != nil {
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: errors.Wrap(err, "failed to send msg").Error(),
		}

	}

	return c.JSON(&fiber.Map{"status": "success"})
}
