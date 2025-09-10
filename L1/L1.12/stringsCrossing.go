package main

import "fmt"

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}

	res := make(map[string]struct{})

	for _, v := range strs {
		res[v] = struct{}{}
	}
	fmt.Println(res)
}
