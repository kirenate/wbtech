package services

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"main.go/repositories"
	"main.go/utils"
)

type Consumer struct {
	reader     *kafka.Reader
	repository *repositories.Repository
}

func NewConsumer(reader *kafka.Reader, repository *repositories.Repository) *Consumer {
	return &Consumer{reader: reader, repository: repository}
}

func (r *Consumer) Consume(ctx context.Context) (string, *repositories.Model, error) {

	msg, err := r.reader.FetchMessage(ctx)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to fetch msg")
	}
	myMsg := msg.Value
	var jsonMsg *repositories.Model
	err = json.Unmarshal(myMsg, &jsonMsg)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to unmarshal msg value")
	}

	jsonMsg.Delivery.OrderUID = jsonMsg.Order.OrderUID
	jsonMsg.Delivery.ID = uuid.New()
	jsonMsg.Payment.OrderUID = jsonMsg.Order.OrderUID
	jsonMsg.Payment.ID = uuid.New()
	items := *jsonMsg.Item
	for i := range items {
		items[i].OrderUID = jsonMsg.Order.OrderUID
		items[i].ID = uuid.New()
	}
	jsonMsg.Item = &items

	err = r.repository.CreateOrderTX(ctx, jsonMsg)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to create order")
	}

	err = r.reader.CommitMessages(ctx, msg)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to commit msg")
	}
	return string(myMsg), jsonMsg, nil
}

func (r *Consumer) Connect(ctx context.Context) (*kafka.Message, error) {

	conn, err := kafka.DialLeader(ctx, "tcp", utils.MyConfig.Kafka, "events", 0)
	if err != nil {
		panic(errors.Wrap(err, "failed to dial kafka"))
	}
	msg, err := conn.ReadMessage(1e6)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read msg")
	}

	return &msg, nil
}
