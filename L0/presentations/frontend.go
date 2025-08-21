package presentations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"io"
	"os"
)

func (r *Presentation) getOrderFront(c *fiber.Ctx) error {
	file, err := os.Open(".data/frontend.html")
	if err != nil {
		return errors.Wrap(err, "failed to open frontend")
	}
	contents, err := io.ReadAll(file)
	if err != nil {
		return errors.Wrap(err, "failed to read frontend")
	}

	return c.Type("html").Send(contents)
}
