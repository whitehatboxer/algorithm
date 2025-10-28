package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func trap(height []int) int {
	res := 0

	stack := []int{}
	for i, n := range height {
		for len(stack) > 0 && n > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}

			left := stack[len(stack)-1]
			w := i - left - 1
			h := min(height[left], n) - height[top]

			res += w * h
		}

		stack = append(stack, i)
	}

	return res
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))

	fmt.Println(trap([]int{4, 2, 3, 2, 3, 6}))
}
