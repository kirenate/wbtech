package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := generator()
	for range 10 {
		wg.Add(1)
		go worker(ch, &wg)
	}
	wg.Wait()
}

func generator() chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for {
			v := rand.Int()
			if v%10 >= 8 {
				fmt.Println("value >= 8")
				return
			}
			ch <- v
		}
	}()

	return ch
}

func worker(in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		v := <-in
		if v%2 != 0 || v == 0 {
			fmt.Println("value is not even")
			return
		}
		fmt.Println(v)
	}
}
