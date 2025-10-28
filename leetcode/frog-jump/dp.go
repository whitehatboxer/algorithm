package main

func canCrossFromDP(stones []int) bool {
	if len(stones) <= 1 {
		return true
	}

	if stones[1] != 1 {
		return false
	}

	stoneMap := make(map[int]int, len(stones))
	for i := 0; i < len(stones); i++ {
		stoneMap[stones[i]] = i
	}

	dp := make([][]bool, len(stones))
	for i := 0; i < len(stones); i++ {
		dp[i] = make([]bool, len(stones)+1)
	}

	dp[1][1] = true

	for i := 1; i < len(stones); i++ {
		for k := 0; k < len(dp[i]); k++ {
			if !dp[i][k] {
				continue
			}

			for s := -1; s <= 1; s++ {
				nextStep := k + s

				if nextStep <= 0 {
					continue
				}
				nextStone := stones[i] + nextStep
				if j, ok := stoneMap[nextStone]; ok {
					dp[j][nextStep] = true
				}
			}
		}

	}

	for _, k := range dp[len(stones)-1] {
		if k {
			return true
		}
	}

	return false
}
