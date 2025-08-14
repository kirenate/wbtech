package services

import (
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
