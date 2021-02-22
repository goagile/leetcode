package fast

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	old, new := x, 0
	for ; x != 0; x /= 10 {
		new = new*10 + x%10
	}
	return old == new
}
