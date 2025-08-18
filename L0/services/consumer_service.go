package services

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"main.go/repositories"
)

func (r *Service) BackgroundConsumer(ctx context.Context) {
	for {
		go func() {
			msg, err := r.reader.FetchMessage(ctx)
			if err != nil {
				defer func() {
					if r := recover(); r != nil {
						log.Error().Err(err).Msg("failed to fetch msg, recovered")
					}
				}()
			}
			myMsg := msg.Value
			var jsonMsg *repositories.Model
			err = json.Unmarshal(myMsg, &jsonMsg)
			if err != nil {
				defer func() {
					if r := recover(); r != nil {
						log.Error().Err(err).Msg("failed to unmarshal msg value, recovered")
					}
				}()
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
				defer func() {
					if r := recover(); r != nil {
						log.Error().Err(err).Msg("failed to create order, recovered")
					}
				}()
			}
			err = r.reader.CommitMessages(ctx, msg)
			if err != nil {
				defer func() {
					if r := recover(); r != nil {
						log.Error().Err(err).Msg("failed to commit msg, recovered")
					}
				}()
			}
		}()
	}
}
