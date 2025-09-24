package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 36, 7, 8, 1, -22, -3, -6, 7, 8}
	nums = quicksort(nums)
	fmt.Println(nums)
}

func quicksort(nums []int) []int {
	var res []int
	nums = recursiveSort(nums, res)
	return nums
}

func recursiveSort(nums, res []int) []int {
	if len(nums) == 1 {
		res = append(res, nums[0])
		return res
	}
	minim := nums[0]
	minIdx := 0
	for i, v := range nums {
		if v < minim {
			minim = v
			minIdx = i
		}
	}
	tmp := nums[0]
	nums[0] = minim
	nums[minIdx] = tmp
	res = append(res, minim)
	if len(nums) >= 1 {
		res = recursiveSort(nums[1:], res)

	}
	return res
}
