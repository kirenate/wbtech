package services

import (
	"context"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"main.go/repositories"
)

type Service struct {
	repository *repositories.Repository
	reader     *kafka.Reader
}

func NewService(repository *repositories.Repository, reader *kafka.Reader) *Service {
	service := &Service{repository: repository, reader: reader}
	service.BackgroundConsumer(context.Background())
	return service
}

func (r *Service) GetOrder(ctx context.Context, orderUID string) (*repositories.Model, error) {
	order, err := r.repository.GetOrderTX(ctx, orderUID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get order from db")
	}

	return order, nil
}
