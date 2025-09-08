package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	after := time.After(1 * time.Second)
	done := make(chan interface{})
	in := generator(done)
	multiply(in, done)

	select {
	case <-after:
		close(done)
		fmt.Println("Exiting...")
		return
	}

}

func generator(done <-chan interface{}) chan int {
	nums := make(chan int, len(os.Args)-1)
	go func() {
		for {
			select {
			case <-done:
				close(nums)
				fmt.Println("nums closed")
				return
			default:
				v := rand.Int()
				nums <- v
				fmt.Println("sent to nums: ", v)
			}
		}
	}()
	return nums
}

func multiply(in <-chan int, done <-chan interface{}) {
	go func() {
		for {
			select {
			case <-done:
				return
			case v := <-in:
				fmt.Println(2 * v)
			}
		}
	}()
}
