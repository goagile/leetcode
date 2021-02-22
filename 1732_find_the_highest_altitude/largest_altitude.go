package largest

import "math"

func largestAltitude(gain []int) int {
	var max, alt int
	for _, g := range gain {
		alt += g
		if alt > max {
			max = alt
		}
	}
	return max
}

func largestAltitudeMathMax(gain []int) int {
	var (
		max float64
		alt int
	)
	for _, g := range gain {
		alt += g
		max = math.Max(max, float64(alt))
	}
	return int(max)
}
