package services

import (
	"context"
	"crypto/rand"
	_ "embed"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"main.go/repositories"
	"main.go/utils"
)

//go:embed model.json
var baseMsg []byte

type Producer struct {
	writer *kafka.Writer
}

func NewProducer() *Producer {
	writer := kafka.Writer{
		Addr:     kafka.TCP(utils.MyConfig.Kafka),
		Topic:    "events",
		Balancer: &kafka.LeastBytes{},
	}

	return &Producer{writer: &writer}
}

func (r *Producer) generateMsg() (*[]byte, error) {
	var model *repositories.Model
	err := json.Unmarshal(baseMsg, &model)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal base json model into model")
	}

	model.Order.OrderUID = rand.Text()

	msg, err := json.Marshal(model)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal model into json")
	}

	return &msg, nil
}

func (r *Producer) SendMsg(ctx context.Context) error {
	for range 10 {
		msg, err := r.generateMsg()
		if err != nil {
			return errors.Wrap(err, "failed to generate msg")
		}

		kafkaMsg := kafka.Message{Value: *msg}

		err = r.writer.WriteMessages(ctx, kafkaMsg)
		if err != nil {
			return errors.Wrap(err, "failed to write msg into kafka")
		}
	}

	return nil
}
