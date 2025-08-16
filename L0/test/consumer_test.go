package test

import (
	"context"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"main.go/repositories"
	"main.go/services"
	"main.go/utils"
	"testing"
)

func TestConsumer(t *testing.T) {
	r := createService()
	err := r.Consumer(context.Background())
	if err != nil {
		assert.IsEqual(err.Error(), "EOF")
	}
}

func createService() *services.Service {
	err := utils.NewConfig("../.data/dev.yaml")
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		utils.MyConfig.Host, utils.MyConfig.Port, utils.MyConfig.User, utils.MyConfig.Password, utils.MyConfig.DBName)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(errors.Wrap(err, "failed to connect database"))
	}

	err = db.AutoMigrate(&repositories.Order{}, &repositories.Payment{}, &repositories.Delivery{}, &repositories.Item{})
	if err != nil {
		panic(errors.Wrap(err, "failed to merge database"))
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{utils.MyConfig.Kafka},
		Topic:   "events",
	})
	defer reader.Close()

	repository := repositories.NewRepository(db)

	service := services.NewService(repository, reader)

	return service
}
