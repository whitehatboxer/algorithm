package main

import "fmt"

/**
 * Input: s = "ADOBECODEBANC", t = "ABC"
 * Output: "BANC"
 * Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
 */

func minWindow(s string, t string) string {
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println("res is: %s", minWindow(s, t))
}
