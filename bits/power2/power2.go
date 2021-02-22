package main

import "fmt"

func main() {
	for i := 0; i < 129; i++ {
		if isPowerOfTwo(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
