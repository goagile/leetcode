package xor

func xorOperation(n int, start int) int {
	x := 0
	for i := 0; i < n; i++ {
		x ^= start + 2*i
	}
	return x
}
