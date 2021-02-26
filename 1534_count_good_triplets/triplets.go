package triplets

import (
	"math"
)

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var z int
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if int(math.Abs(float64(arr[i]-arr[j]))) <= a &&
					int(math.Abs(float64(arr[j]-arr[k]))) <= b &&
					int(math.Abs(float64(arr[i]-arr[k]))) <= c {
					z++
				}
			}
		}
	}
	return z
}

func countGoodTriplets2(arr []int, a int, b int, c int) int {
	var counter int
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			// don't need to check arr[j+1:k] if first condition false
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[j]-arr[k]) <= b &&
					abs(arr[i]-arr[k]) <= c {
					counter++
				}
			}
		}
	}
	return counter
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
