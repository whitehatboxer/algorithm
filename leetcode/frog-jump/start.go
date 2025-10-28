package main

func canCrossFromStart(stones []int) bool {
	if len(stones) == 0 {
		return false
	}

	sMap := map[int]int{}
	for i := 0; i < len(stones); i++ {
		sMap[stones[i]] = i
	}

	if _, ok := sMap[1]; !ok {
		return false
	}

	return searchFromStart(sMap, stones, len(stones), 1, 1)

}

func searchFromStart(sMap map[int]int, stones []int, length, index, step int) bool {
	// fmt.Println("index is: ", index, length)
	if index == length-1 {
		return true
	}

	for j := -1; j <= 1; j++ {
		new_step := step + j
		if new_step == 0 {
			continue
		}
		next_stone := stones[index] + new_step
		if v, ok := sMap[next_stone]; ok {
			// scan next
			if searchFromStart(sMap, stones, length, v, new_step) == true {
				return true
			}
		}
	}

	return false
}
