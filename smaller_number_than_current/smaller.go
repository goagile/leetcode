package smaller

func smallerNumbersThanCurrent(nums []int) []int {
	z := make([]int, len(nums))
	for i, a := range nums {
		for _, b := range nums {
			if b < a {
				z[i]++
			}
		}
	}
	return z
}
