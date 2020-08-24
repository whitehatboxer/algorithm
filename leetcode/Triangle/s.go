package main

import (
	"fmt"
)

func min(v1 int, v2 int) int {
	if v1 > v2 {
		return v2
	} else {
		return v1
	}
}

func minimumTotal(triangle [][]int) int {
	var height int
	height = len(triangle)

	for i := height - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}

func main() {
	a := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println("res is: ", minimumTotal(a))
}
