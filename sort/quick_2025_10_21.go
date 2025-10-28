package main

import "fmt"

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	target := nums[start]
	i, j := start, end

	for i < j {
		for i < j && nums[j] >= target {
			j--
		}

		for i < j && nums[i] <= target {
			i++
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[j], nums[start] = nums[start], nums[j]

	quickSort(nums, start, j-1)
	quickSort(nums, j+1, end)

}

func main() {
	nums := []int{1, 1, -4, -4, 5, 9}
	// nums := []int{1, 1, 1}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
