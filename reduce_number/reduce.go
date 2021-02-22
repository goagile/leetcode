package reduce

func numberOfSteps(x int) int {
	c := 0
	for x > 0 {
		if x%2 == 0 {
			x /= 2
		} else {
			x--
		}
		c++
	}
	return c
}
