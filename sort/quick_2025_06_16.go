package main

import "fmt"

func quickSort(array []int, start, end int) {
	if start >= end {
		return
	}

	target := array[start]
	i, j := start, end

	for i < j {
		for i < j && array[j] >= target {
			j--
		}

		for i < j && array[i] <= target {
			i++
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	array[j], array[start] = array[start], array[j]

	quickSort(array, start, j-1)
	quickSort(array, j+1, end)
}

func main() {
	nums := []int{1, 1, -4, -4, 5, 9}
	// nums := []int{1, 1, 1}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
