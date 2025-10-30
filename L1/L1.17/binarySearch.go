package main

import (
	"fmt"
	"slices"
)

func main() {
	exp := []int{1, 2, 3, 4, 5, 6, 8, 8, 9, 7}
	slices.Sort(exp) //использую чтобы не париться с условием сортировки входных данных

	res := binarySearch(exp, 7)
	fmt.Println(res)
}

func binarySearch(nums []int, target int) int {
	idx := len(nums) / 2
	elem := nums[idx]

	if elem == target {
		return idx
	}
	for elem != target {
		if len(nums) <= 1 {
			return -1
		}
		if elem < target {
			nums = nums[len(nums)/2+1:]
			idx += len(nums) / 2
			elem = nums[len(nums)/2]
			continue
		}
		nums = nums[0 : len(nums)/2]
		idx -= len(nums) / 2
		elem = nums[len(nums)/2]
	}
	return idx
}
