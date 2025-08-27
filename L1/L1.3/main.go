package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
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
	i := 0
	for range workers {
		i++
		wg.Add(1)
		go worker(ch, &wg, i)
	}
	for {
		ch <- rand.Int()
	}
	wg.Wait()

}

func worker(in chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for {
		fmt.Println("worker ", i)
		fmt.Println(<-in)
	}

}
