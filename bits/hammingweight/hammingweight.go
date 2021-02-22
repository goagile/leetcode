package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingWeight(0b00000000000000000000000000001011))
	fmt.Println(hammingWeight(0b00000000000000000000000010000000))
	fmt.Println(hammingWeight(0b11111111111111111111111111111101))
}

var cache = make(map[uint32]int)

func hammingWeight(num uint32) int {
	if v, ok := cache[num]; ok {
		return v
	}
	c := 0
	for _, x := range fmt.Sprintf("%b", num) {
		if x == '1' {
			c++
		}
	}
	cache[num] = c
	return c
}
