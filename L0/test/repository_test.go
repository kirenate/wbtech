package test

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetOrderTX(t *testing.T) {
	repository := createRepository(t)
	order, err := repository.GetOrderTX(context.Background(), "WQLMV3NQ2CQYCP4LNJDQGDZLTB")
	require.NoError(t, err)
	require.NotEmpty(t, order)

	fmt.Println(order.Order)
	fmt.Println(order.Delivery)
	fmt.Println(order.Payment)
	fmt.Println(order.Item)
}

func TestGetOrderTXFail(t *testing.T) {
	repository := createRepository(t)
	order, err := repository.GetOrderTX(context.Background(), "not-real-uid")
	require.NotEmpty(t, err)

	errors.Is(err, errors.New("failed to find order model: order does not exist"))

	require.Empty(t, order)
}
