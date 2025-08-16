package test

import (
	"context"
	"github.com/stretchr/testify/require"
	"main.go/utils"
	"testing"
)

func TestSendMsg(t *testing.T) {
	ctx := context.Background()

	err := utils.NewConfig("../.data/.yaml")
	require.NoError(t, err)

	producer := utils.NewProducer()
	err = producer.SendMsg(ctx)
	require.NoError(t, err)
}
