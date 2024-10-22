package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	triples := make([][]int, 0)
	target := 0

	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				triples = append(triples, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > target {
				k--
			} else {
				j++
			}
		}
	}

	return triples
}

func main() {
	testList := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(testList))

	testList = []int{-2, 0, 1, 1, 2}
	fmt.Println(threeSum(testList))
}
