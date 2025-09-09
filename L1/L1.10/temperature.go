package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 20}
	slices.Sort(temps)
	idx := math.Round(temps[0]/10)*10 + 10
	res := make(map[int][]float64)
	for i := 0; i < len(temps); i++ {
		if idx < 0 {
			if temps[i] <= idx {
				res[int(idx)] = append(res[int(idx)], temps[i])
			} else {
				idx += 10
				i--
			}
		} else {
			if temps[i]-10 <= idx {
				res[int(idx)] = append(res[int(idx)], temps[i])
			} else {
				idx += 10
				i--
			}
		}

	}
	fmt.Println(res)
}
