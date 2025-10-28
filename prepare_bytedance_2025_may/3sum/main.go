package main

import "fmt"

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	target := nums[start]
	i, j := start, end
	for j > i {
		for j > i && nums[j] >= target {
			j--
		}

		for j > i && nums[i] <= target {
			i++
		}

		if j > i {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[j], nums[start] = nums[start], nums[j]
	quickSort(nums, start, j-1)
	quickSort(nums, j+1, end)
}

func threeSum(nums []int) [][]int {
	length := len(nums)
	quickSort(nums, 0, length-1)
	fmt.Println(nums)

	res := [][]int{}
	for i := 0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := length - 1
		lastJ, lastK := nums[0]-1, nums[0]-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum == 0 {
				if nums[j] != lastJ || nums[k] != lastK {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					lastJ = nums[j]
					lastK = nums[k]
				}
				j++
			} else {
				k--
			}

		}
	}

	return res
}

func threeSumWrong(nums []int) [][]int {
	length := len(nums)
	res := [][]int{}
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length-1; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)

	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{1, 1, 1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
	fmt.Println(threeSum([]int{-2, 0, 0, 0, 2}))
}
