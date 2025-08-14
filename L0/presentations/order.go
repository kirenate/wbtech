package presentations

import (
	"github.com/segmentio/kafka-go"
	"main.go/services"
)

type Presentation struct {
	service *services.Service
	writer  *kafka.Writer
}

func NewPresentation(service *services.Service, writer *kafka.Writer) *Presentation {
	return &Presentation{service: service, writer: writer}
}
