package slow

import "fmt"

func IsPalindrome(x int) bool {
	s := fmt.Sprintf("%v", x)
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
