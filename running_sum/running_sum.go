package runningsum

func runningSum(s []int) []int {
	x, n := 0, len(s)
	z := make([]int, n)
	for i := 0; i < n; i++ {
		x += s[i]
		z[i] = x
	}
	return z
}

func runningSumShort(s []int) []int {
	for i := 1; i < len(s); i++ {
		s[i] += s[i-1]
	}
	return s
}
