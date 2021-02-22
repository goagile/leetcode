package ps

func subtractProductAndSum(n int) int {
	s, p := 0, 1
	for ; n != 0; n /= 10 {
		r := (n % 10)
		p *= r
		s += r
	}
	return (p - s)
}
