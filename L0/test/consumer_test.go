package test

import (
	"fmt"
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

func createRepository(t *testing.T) *repositories.Repository {
	t.Helper()
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
	repository := repositories.NewRepository(db)
	return repository
}

func createService(t *testing.T) *services.Service {
	t.Helper()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{utils.MyConfig.Kafka},
		GroupID:     "0",
		Topic:       "events",
		Partition:   0,
		MaxAttempts: 5,
	})
	repository := createRepository(t)
	consumer := services.NewService(repository, reader)

	return consumer
}

func createConfig(t *testing.T) {
	t.Helper()

	err := utils.NewConfig("../.data/dev.yaml")
	if err != nil {
		panic(err)
	}
}
