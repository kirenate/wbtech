package test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetOrderTX(t *testing.T) {
	repository := createRepository(t)
	order, err := repository.GetOrderTX(context.Background(), "WQLMV3NQ2CQYCP4LNJDQGDZLTB")
	require.NoError(t, err)
	require.NotEmpty(t, order)

	fmt.Println(order.OrderUID)
	fmt.Println(order.Order)
	fmt.Println(order.Delivery)
	fmt.Println(order.Payment)
	fmt.Println(order.Item)
}
