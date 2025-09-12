package main

import "fmt"

func main() {
	in := gen()
	detect(in)
}

func gen() chan interface{} {
	out := make(chan interface{}, 4)
	var a string
	var b int
	var c bool
	var d chan interface{}
	out <- a
	out <- b
	out <- c
	out <- d
	close(out)
	return out
}

func detect(in chan interface{}) {
	for {
		select {
		case v, ok := <-in:
			if !ok {
				return
			}
			fmt.Printf("%T\n", v)
		}
	}

}
