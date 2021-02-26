package even

func findNumbers(nums []int) int {
	evens := 0
	for _, num := range nums {
		digits := 0
		for n := num; n > 0; n /= 10 {
			digits++
		}
		if 0 == digits%2 {
			evens++
		}
	}
	return evens
}

func findNumbers2(nums []int) int {
	evens := 0
	for i := 0; i < len(nums); i++ {
		digits := 1
		for n := nums[i] / 10; n != 0; n /= 10 {
			digits++
		}
		if 0 == digits%2 {
			evens++
		}
	}
	return evens
}
