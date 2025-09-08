package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("Please enter at least 2 integers"))
	}
	num, ibit := os.Args[1], os.Args[2]
	res, err := strconv.Atoi(num)
	if err != nil {
		panic(errors.Wrap(err, "Please, enter valid integer"))
	}
	i, err := strconv.Atoi(ibit)
	if err != nil {
		panic(errors.Wrap(err, "Please, enter valid integer"))
	}
	res ^= 1 << i
	fmt.Println(res)
}
