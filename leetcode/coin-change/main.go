package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return -1
	}

	for _, c := range coins {
		if c == amount {
			return 1
		}
	}

	var res int
	for _, c := range coins {
		temp := coinChange(coins, amount-c) + 1
		if res > temp && temp != -1 {
			res = temp
		}
	}
	return res
}

func main() {
	coins := []int{1, 3, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
