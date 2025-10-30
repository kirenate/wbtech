package main

import (
	"fmt"
	"sync"
)

type Val struct {
	Value int64
	wg    sync.WaitGroup
	mu    sync.Mutex
}

func (r *Val) Increment() {
	r.mu.Lock()
	r.Value++
	r.mu.Unlock()
	r.wg.Done()
}

func main() {
	val := Val{Value: 1, mu: sync.Mutex{}, wg: sync.WaitGroup{}}

	for range 10000 {
		val.wg.Add(1)
		go val.Increment()
	}
	val.wg.Wait()
	fmt.Println(val.Value)
}
