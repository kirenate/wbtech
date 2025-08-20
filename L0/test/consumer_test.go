package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"main.go/repositories"
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
		Logger: logger.Default.LogMode(logger.Error), NamingStrategy: schema.NamingStrategy{SingularTable: true}})
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

type TestService struct {
	Repository  *repositories.Repository
	Reader      *kafka.Reader
	RedisClient *redis.Client
}

func createTestService(t *testing.T) *TestService {
	t.Helper()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{utils.MyConfig.Kafka},
		GroupID:     "0",
		Topic:       "events",
		Partition:   0,
		MaxAttempts: 5,
	})
	repository := createRepository(t)
	redisClient := createRedis(t)

	return &TestService{Repository: repository, Reader: reader, RedisClient: redisClient}
}

func createConfig(t *testing.T) {
	t.Helper()

	err := utils.NewConfig("../.data/dev.yaml")
	if err != nil {
		panic(err)
	}
}

func TestBackgroundConsumer(t *testing.T) {
	BackgroundConsumer(t, context.Background())
}

func BackgroundConsumer(t *testing.T, ctx context.Context) {
	createConfig(t)
	r := createTestService(t)
	msg, err := r.Reader.FetchMessage(ctx)
	if err != nil {
		if r := recover(); r != nil {
			log.Error().Err(err).Msg("failed to fetch msg, recovered")
		}
	}

	fmt.Println("fetched msg: ", string(msg.Value))

	myMsg := msg.Value
	var jsonMsg *repositories.Model
	err = json.Unmarshal(myMsg, &jsonMsg)
	if err != nil {
		if r := recover(); r != nil {
			log.Error().Err(err).Msg("failed to unmarshal msg value, recovered")
		}
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

	fmt.Println("jsonMsg: ", jsonMsg)
	err = r.Repository.CreateOrderTX(ctx, jsonMsg)
	if err != nil {
		if r := recover(); r != nil {
			log.Error().Err(err).Msg("failed to create order, recovered")
		}
	}

	fmt.Println("order created")
	res, err := json.Marshal(jsonMsg)
	require.NoError(t, err)

	stat := r.RedisClient.Set(ctx, jsonMsg.Order.OrderUID, res, utils.RedisCFG.TTL)
	if stat.Err() != nil {
		if r := recover(); r != nil {
			log.Error().Err(stat.Err()).Msg("failed to set msg in cache, recovered")
		}
	}

	fmt.Println("\n\nredis response: ", stat)

	err = r.Reader.CommitMessages(ctx, msg)
	if err != nil {
		if r := recover(); r != nil {
			log.Error().Err(err).Msg("failed to commit msg, recovered")
		}
	}
	fmt.Println("msg commited")
}
