package xored

func decode(encoded []int, first int) []int {
	z := []int{first}
	for _, e := range encoded {
		first ^= e
		z = append(z, first)
	}
	return z
}
