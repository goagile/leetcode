package discount

func finalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			discount := prices[j]
			if discount <= prices[i] {
				prices[i] -= discount
				break
			}
		}
	}
	return prices
}
