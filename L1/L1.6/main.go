package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ch := generator(ctx)
	after := time.After(1 * time.Second)
	var wg sync.WaitGroup
	done := make(chan interface{})
	go workerCtx(ctx, ch)
	go workerChan(done, ch)
	wg.Add(1)
	go workerWg(ch, &wg)
	select {
	case <-after:
		ctx.Done()
		close(done)
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

func workerCtx(ctx context.Context, in chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("ctx worker: ", <-in)
		}
	}
}

func workerChan(done chan interface{}, in chan int) {
	for {
		select {
		case <-done:
			return
		case v := <-in:
			fmt.Println("channel worker: ", v)
		}
	}
}

func workerWg(in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		v := <-in
		if v < 14946641 {
			fmt.Println("value is small")
			runtime.Goexit()
		}
		fmt.Println("wg worker: ", v)
	}
}
