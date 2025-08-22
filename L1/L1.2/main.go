package main

import (
	"fmt"
)

func main() {
	num := [5]int{2, 4, 6, 8, 10}

	ch := gen(num)
	out := calc(ch)

	for n := range out {
		fmt.Println(n)
	}

}

func gen(sl [5]int) chan int {
	ch := make(chan int)
	go func() {
		for _, v := range sl {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func calc(ch chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			out <- n * n
		}
		close(out)
	}()
	return out
}
