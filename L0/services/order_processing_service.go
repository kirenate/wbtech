package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"main.go/repositories"
)

type Service struct {
	repository  *repositories.Repository
	reader      *kafka.Reader
	redisClient *redis.Client
}

func NewService(repository *repositories.Repository, reader *kafka.Reader, redisClient *redis.Client) *Service {
	service := &Service{repository: repository, reader: reader, redisClient: redisClient}
	service.BackgroundConsumer(context.Background())
	return service
}

func (r *Service) GetOrder(ctx context.Context, orderUID string) (*repositories.Model, error) {
	var order *repositories.Model
	stat := r.redisClient.Get(ctx, orderUID)
	if errors.Is(stat.Err(), redis.Nil) {
		log.Info().Interface("stat", stat).Send()
		order, err := r.repository.GetOrderTX(ctx, orderUID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get order from db")
		}
		return order, nil
	}
	err := stat.Scan(order)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan into order")
	}

	return order, nil
}
