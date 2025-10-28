package main

import "fmt"

func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func main() {
	nums := []int{1, 3, -5, -5, -5, 9, 0}

	fmt.Println(twoSum(nums, 12))
}
