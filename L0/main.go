package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"main.go/presentations"
	"main.go/repositories"
	"main.go/services"
	"time"

	"main.go/utils"
)

var log = utils.MakeLogger()

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("logger.started")

	err := utils.NewConfig(".data/dev.yaml")
	if err != nil {
		panic(errors.Wrap(err, "failed to create config"))
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		utils.MyConfig.Host, utils.MyConfig.Port, utils.MyConfig.User, utils.MyConfig.Password, utils.MyConfig.DBName)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	dbPool, err := db.DB()
	if err != nil {
		panic(errors.Wrap(err, "failed to set database connection pool"))
	}

	dbPool.SetMaxIdleConns(utils.MyConfig.MaxIdleConns)
	dbPool.SetMaxOpenConns(utils.MyConfig.MaxOpenConns)
	dbPool.SetConnMaxLifetime(utils.MyConfig.ConnMaxLifetimeMinutes * time.Minute)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect database"))
	}

	err = db.AutoMigrate(&repositories.Order{}, &repositories.Payment{}, &repositories.Delivery{}, &repositories.Item{})
	if err != nil {
		panic(errors.Wrap(err, "failed to merge database"))
	}

	log.Info().Msgf("db.%s.started.at %s:%d", utils.MyConfig.DBName, utils.MyConfig.Host, utils.MyConfig.Port)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{utils.MyConfig.Kafka},
		Topic:   "events",
	})
	defer reader.Close()

	repository := repositories.NewRepository(db)

	service := services.NewService(repository, reader)

	presentation := presentations.NewPresentation(service)

	app := presentation.BuildApp()

	err = app.Listen(utils.MyConfig.Addr)
	if err != nil {
		panic(errors.Wrap(err, "failed to start server"))
	}

}
