package main

import "fmt"

func main() {
	var1 := -2
	var2 := 500
	fmt.Println(var1, var2)
	//
	var1 ^= var2
	var2 ^= var1
	var1 ^= var2
	//
	fmt.Println(var1, var2)
}
