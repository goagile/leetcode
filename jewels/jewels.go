package jewels

import (
	"strings"
)

func numJewelsInStones(jewels string, stones string) int {
	c := 0
	for _, j := range jewels {
		for _, s := range stones {
			if j == s {
				c++
			}
		}
	}
	return c
}

func numJewelsInStonesContains(jewels string, stones string) int {
	c := 0
	for _, s := range stones {
		if strings.Contains(jewels, string(s)) {
			c++
		}
	}
	return c
}

func numJewelsInStonesHashMap(jewels string, stones string) int {
	jmap := make(map[rune]bool)
	for _, j := range jewels {
		jmap[j] = true
	}
	c := 0
	for _, s := range stones {
		if _, ok := jmap[s]; ok {
			c++
		}
	}
	return c
}
