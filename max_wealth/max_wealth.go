package maxwealth

func maximumWealthRange(accounts [][]int) int {
	max := 0
	for _, banks := range accounts {
		sum := 0
		for _, bank := range banks {
			sum += bank
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func maximumWealthLoop(accounts [][]int) int {
	max := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
