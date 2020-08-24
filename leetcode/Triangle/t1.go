package main

import "fmt"

func main() {
	var t1, t2 [][]int
	t1 = [][]int{
		{1},
		{2, 3},
	}

	copy(t2, t1)
	fmt.Println(t1)
	fmt.Println(t2)
}
