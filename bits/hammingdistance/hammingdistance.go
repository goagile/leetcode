package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	return hammingWeight(x ^ y)
}

func hammingWeight(z int) int {
	c := 0
	for _, x := range fmt.Sprintf("%b", z) {
		if '1' == x {
			c++
		}
	}
	return c
}
