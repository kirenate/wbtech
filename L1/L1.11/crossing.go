package main

import (
	"fmt"
	"slices"
)

func main() {
	first := []int{1, 20, 3, 40}
	second := []int{20, 3, 40}
	var res []int
	for _, v := range first {
		if slices.Contains(second, v) {
			res = append(res, v)
		}
	}

	fmt.Println(res)
}
