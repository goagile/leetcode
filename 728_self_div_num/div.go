package div

func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for n := left; n <= right; n++ {
		if isSelfDiveded(n) {
			result = append(result, n)
		}
	}
	return result
}

func isSelfDiveded(n int) bool {
	for i := n; i > 0; i /= 10 {
		x := i % 10
		if x == 0 || n%x != 0 {
			return false
		}
	}
	return true
}
