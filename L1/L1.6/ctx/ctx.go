package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx := context.Background()
	ch := generator(ctx)
	after := time.After(2 * time.Second)
	for range 10 {
		go worker(ctx, ch)
	}
	select {
	case <-after:
		ctx.Done()
	}
}

func generator(ctx context.Context) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				out <- rand.Int()
			}
		}
	}()
	return out

}

func worker(ctx context.Context, in chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(<-in)
		}
	}
}
