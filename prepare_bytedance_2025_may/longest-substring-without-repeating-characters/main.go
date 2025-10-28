package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	charMap := map[byte]int{}

	windowStart := 0
	windowEnd := 0
	maxWindow := 0
	for {
		if windowEnd >= length {
			break
		}

		currentChar := s[windowEnd]
		if lastPos, ok := charMap[currentChar]; ok && lastPos >= windowStart {
			windowStart = charMap[currentChar] + 1
		}
		charMap[currentChar] = windowEnd

		if windowEnd-windowStart+1 > maxWindow {
			maxWindow = windowEnd - windowStart + 1
		}
		windowEnd += 1
	}

	return maxWindow
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}
