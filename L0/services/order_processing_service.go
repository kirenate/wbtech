package services

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"main.go/repositories"
)

type Service struct {
	repository *repositories.Repository
	reader     *kafka.Reader
}

func NewService(repository *repositories.Repository, reader *kafka.Reader) *Service {
	return &Service{repository: repository, reader: reader}
}

var Validate = validator.New(validator.WithRequiredStructEnabled())

func (r *Service) ProcessMessage(ctx context.Context) error {
	msg, err := r.reader.FetchMessage(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to fetch msg from kafka")
	}

	var req *repositories.Data
	err = json.Unmarshal(msg.Value, &req)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal msg")
	}

	err = Validate.Struct(&req)
	if err != nil {
		return errors.Wrap(err, "failed to validate order request")
	}

	err = r.repository.CreateOrderTX(ctx, req)
	if err != nil {
		return errors.Wrap(err, "failed to create order")
	}

	return nil
}

func (r *Service) GetOrder(ctx context.Context, orderUID string) (*repositories.Data, error) {
	order, err := r.repository.GetOrderTX(ctx, orderUID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get order from db")
	}

	return order, nil
}
