package main

import "fmt"

func main() {
	// stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	stones := []int{0, 1, 4, 5, 6, 8, 12, 17}

	// fmt.Println(canCrossFromStart(stones))
	fmt.Println(canCrossFromDP(stones))
}
