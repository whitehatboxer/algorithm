package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	vMap := make(map[int]int, 0)
	res := make([]int, len(nums1))
	stack := make([]int, 0)

	for _, n := range nums2 {
		if len(stack) == 0 {
			stack = append(stack, n)
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] < n {
			v := stack[len(stack)-1]
			vMap[v] = n

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, n)
	}

	for i, n := range nums1 {
		if v, ok := vMap[n]; !ok {
			res[i] = -1
		} else {
			res[i] = v
		}
	}

	return res
}

// Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
// Output: [-1,3,-1]
// Input: nums1 = [2,4], nums2 = [1,2,3,4]
// Output: [3,-1]
func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))

	nums1 = []int{4, 1, 2}
	nums2 = []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
