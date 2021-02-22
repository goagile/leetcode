package good

func numIdenticalPairs(a []int) int {
	c := 0
	for i, v := range a {
		for j := i + 1; j < len(a); j++ {
			if v == a[j] {
				c++
			}
		}
	}
	return c
}

func numIdenticalPairsHash(nums []int) int {
	freqMap := make(map[int]int)

	for _, value := range nums {
		freqMap[value]++
	}

	goodPairs := 0
	for _, value := range freqMap {
		goodPairs += (value * (value - 1)) / 2
	}

	return goodPairs
}
