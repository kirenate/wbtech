package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic(errors.New("please enter valid integer number of workers"))
	}
	arg := args[1]
	workers, err := strconv.Atoi(arg)
	if err != nil {
		panic(errors.Wrap(err, "please enter valid integer number of workers"))
	}
	ctx := context.Background()
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	ch := make(chan int)

	for range workers {
		go worker(ctx, ch)
	}

	for {
		ch <- rand.Int()
		select {
		case <-sigint:
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
