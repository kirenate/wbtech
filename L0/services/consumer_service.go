package services

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"main.go/repositories"
	"main.go/utils"
)

func (r *Service) BackgroundConsumer(ctx context.Context) {
	for {
		msg, err := r.reader.FetchMessage(ctx)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Send()
		}
		err = r.consume(ctx, &msg)
		if err == nil {
			err = r.reader.CommitMessages(ctx, msg)
			if err != nil {
				log.Ctx(ctx).Error().Err(err).Send()
			}
		} else {
			log.Ctx(ctx).Error().Err(err).Send()
		}
	}
}

func (r *Service) consume(ctx context.Context, msg *kafka.Message) error {
	myMsg := msg.Value
	var jsonMsg *repositories.Model
	err := json.Unmarshal(myMsg, &jsonMsg)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal msg value")
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
		return errors.Wrap(err, "failed to create order")
	}
	res, err := json.Marshal(jsonMsg)
	if err != nil {
		return errors.Wrap(err, "failed to marshal msg")
	}
	stat := r.redisClient.Set(ctx, jsonMsg.Order.OrderUID, res, utils.RedisCFG.TTL)
	if stat.Err() != nil {
		return errors.Wrap(err, "failed to set msg in cache")
	}
	return nil
}
