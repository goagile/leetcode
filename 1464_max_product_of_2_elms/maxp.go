package maxp

import "sort"

func maxProduct(a []int) int {
	max, n := 0, len(a)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			p := (a[i] - 1) * (a[j] - 1)
			if p > max {
				max = p
			}
		}
	}
	return max
}

func maxProductSort(a []int) int {
	sort.Ints(a)
	return (a[len(a)-1] - 1) * (a[len(a)-2] - 1)
}
