package zero

func sumZero(n int) []int {
	a := make([]int, 0)
	if n == 1 {
		a = append(a, 0)
		return a
	}
	mid := n / 2
	skip := (n%2 == 0)
	for i := -mid; i <= mid; i++ {
		if i == 0 && skip {
			continue
		}
		a = append(a, i)
	}
	return a
}
