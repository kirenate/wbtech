package test

import (
	"context"
	"github.com/stretchr/testify/require"
	"main.go/services"
	"main.go/utils"
	"testing"
)

func TestSendMsg(t *testing.T) {
	ctx := context.Background()

	err := utils.NewConfig("../.data/dev.yaml")
	require.NoError(t, err)

	producer := services.NewProducer()
	err = producer.SendMsg(ctx)
	require.NoError(t, err)
}
