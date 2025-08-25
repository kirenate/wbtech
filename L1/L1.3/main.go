package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"sync"
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
	var wg sync.WaitGroup
	ch := make(chan int)
	for range workers {
		wg.Add(1)
		go worker(ch, &wg)
	}
	for i := range workers {
		ch <- i
	}
	wg.Wait()

}

func worker(in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-in)
}
