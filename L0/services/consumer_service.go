package services

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

func (r *Service) Consumer(ctx context.Context) error {
	msg, err := r.reader.FetchMessage(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to fetch msg")
	}
	fmt.Println("msg: ", msg.Value)
	//var jsonMsg *repositories.Model
	//err = json.Unmarshal(msg.Value, jsonMsg)
	//if err != nil {
	//	return errors.Wrap(err, "failed to unmarshal msg value")
	//}
	//fmt.Println("json msg:", jsonMsg)
	//err = r.repository.CreateOrderTX(ctx, jsonMsg)
	//if err != nil {
	//	return errors.Wrap(err, "failed to create order")
	//}

	return nil
}
