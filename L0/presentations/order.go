package presentations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
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
	if err != nil {
		return errors.Wrap(err, "service failed to get order")
	}

	return c.JSON(order)
}

func (r *Presentation) generateOrders(c *fiber.Ctx) error {
	producer := utils.NewProducer()
	errs := make(chan error)
	for range 10 {
		go func() {
			err := producer.SendMsg(c.UserContext())
			errs <- err
		}()
		if err := <-errs; err != nil {
			return errors.Wrap(err, "failed to send msg")
		}
	}
	
	return nil
}
