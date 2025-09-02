package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan interface{})

	ch := generator(done)
	for range 10 {
		go worker(done, ch)
	}
	time.Sleep(2 * time.Second)
	close(done)

}

func generator(done chan interface{}) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			default:
				ch <- rand.Int()
			}
		}
	}()
	return ch
}

func worker(done chan interface{}, in chan int) {
	for {
		select {
		case <-done:
			return
		case v := <-in:
			fmt.Println(v)
		}
	}
}
