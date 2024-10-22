package main

import (
	"fmt"
)

func quickSort0(nums []int, start, end int) {
	if start >= end {
		return
	}

	target := nums[start]
	i := start
	j := end

	var tmp int
	for {
		if i >= j {
			break
		}

		if nums[j] >= target {
			j--
			continue
		}

		if nums[i] <= target {
			i++
			continue
		}

		tmp = nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
	}

	tmp = nums[i]
	nums[i] = nums[start]
	nums[start] = tmp

	quickSort0(nums, start, i-1)
	quickSort0(nums, i+1, end)
}

func quickSort1(nums []int, start, end int) {
	if start >= end {
		return
	}

	target := nums[start]
	i := start
	j := end

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

	nums[start], nums[i] = nums[i], nums[start]
	quickSort1(nums, start, i-1)
	quickSort1(nums, i+1, end)
}

func quickSort2(nums []int, start, end int) {
	if start >= end {
		return
	}

	// 选择一个中间点作为 target
	mid := (start + end) / 2
	if nums[start] > nums[mid] {
		nums[start], nums[mid] = nums[mid], nums[start]
	}
	if nums[start] > nums[end] {
		nums[end], nums[start] = nums[start], nums[end]
	}
	if nums[mid] > nums[end] {
		nums[end], nums[mid] = nums[mid], nums[end]
	}

	nums[mid], nums[start] = nums[start], nums[mid]

	target := nums[start]
	i := start
	j := end

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

	nums[start], nums[i] = nums[i], nums[start]

	quickSort2(nums, start, i-1)
	quickSort2(nums, i+1, end)

}

func quickSort(nums []int, start, end int) {
	quickSort2(nums, start, end)
}

func sortArray(nums []int) []int {
	length := len(nums)
	quickSort(nums, 0, length-1)
	return nums
}

func main() {
	array := []int{1, 32, 2, -1, 0, 12, 99, 23}
	fmt.Println("origin: ", array)
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)

	array = []int{}
	fmt.Println("origin: ", array)
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)

	array = []int{5, 2, 3, 1}
	fmt.Println("origin: ", array)
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)

	array = []int{110, 100, 0}
	fmt.Println("origin: ", array)
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
