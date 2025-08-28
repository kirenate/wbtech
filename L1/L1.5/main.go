package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic(errors.New("please enter valid integer number of workers"))
	}
	arg := args[1]
	n, err := strconv.Atoi(arg)
	if err != nil {
		panic(errors.Wrap(err, "please enter valid integer number of workers"))
	}

	ch := make(chan int)
	ctx := context.Background()
	after := time.After(time.Duration(n) * time.Second)
	go worker(ctx, ch)

	for {
		ch <- rand.Int()
		select {
		case <-after:
			ctx.Done()
			fmt.Println("Exiting...")
			os.Exit(1)
		default:
			continue
		}

	}

}

func worker(ctx context.Context, in chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-in:
			fmt.Println(<-in)
		default:
			continue
		}
	}

}
