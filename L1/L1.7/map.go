package main

import (
	"fmt"
	"sync"
)

type mmap struct {
	wmap map[int]int
	mu   sync.Mutex
	wg   sync.WaitGroup
}

func main() {
	wmap := mmap{wmap: make(map[int]int), mu: sync.Mutex{}, wg: sync.WaitGroup{}}
	wmap.write(3, 8)
	wmap.write(2, 4)
	wmap.write(1, 2)
	wmap.wg.Wait()
	fmt.Println(wmap.read(3))
}

func (r *mmap) write(addr int, val int) {
	r.wg.Add(1)

	go func() {
		r.mu.Lock()
		defer r.wg.Done()
		defer r.mu.Unlock()
		r.wmap[addr] = val
		return
	}()
	return
}

func (r *mmap) read(addr int) int {
	r.wg.Add(1)

	out := make(chan int)
	go func() {
		r.mu.Lock()
		defer r.wg.Done()
		defer r.mu.Unlock()
		out <- r.wmap[addr]
		close(out)
		return
	}()
	return <-out
}
