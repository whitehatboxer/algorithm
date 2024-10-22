// this function will cause leetcode time limit exceeded
package main

import (
	"fmt"
	"slices"
)

func threeSum0(nums []int) [][]int {
	res := make([][]int, 0)
	keys := map[string]bool{}
	for i, _ := range nums {
		target1 := 0 - nums[i]
		for j, _ := range nums {
			if j >= i {
				continue
			}
			target2 := target1 - nums[j]
			for k, _ := range nums {
				if k >= j || k >= i {
					continue
				}
				if nums[k] == target2 {
					list := []int{nums[i], nums[j], nums[k]}
					slices.Sort(list)
					key := fmt.Sprintf("%d-%d-%d", list[0], list[1], list[2])
					if _, ok := keys[key]; !ok {
						res = append(res, list)
						keys[key] = true
					}
				}
			}
		}
	}
	return res
}
