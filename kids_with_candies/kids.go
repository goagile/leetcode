package kids

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for _, c := range candies {
		if c > max {
			max = c
		}
	}
	z := make([]bool, len(candies))
	for i, c := range candies {
		z[i] = (c + extraCandies) >= max
	}
	return z
}
