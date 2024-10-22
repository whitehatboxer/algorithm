package main

import (
	"fmt"
)

func findDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	res := make([]int, 0)
	for _, n := range nums {
		idx := n - 1
		if n < 0 {
			idx = -n - 1
		}

		if nums[idx] < 0 {
			res = append(res, idx+1)
		} else {
			nums[idx] = -nums[idx]
		}
	}

	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}
