package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestRW(t *testing.T) {
	wmap := mmap{wmap: make(map[int]int), mu: sync.Mutex{}, wg: sync.WaitGroup{}}
	wmap.write(3, 8)
	wmap.write(2, 4)
	wmap.write(1, 2)
	wmap.wg.Wait()
	fmt.Println(wmap.read(3))
	
}
