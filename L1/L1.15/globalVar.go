package main

import (
	"fmt"
)

func someFunc() string {
	v := createHugeString(1 << 10)
	justString := v[:100]
	return justString
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}
