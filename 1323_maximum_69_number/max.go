package maximum

func maximum69Number(num int) int {
	max := num
	tail := 0

	current, remainder := num/10, num%10

	for power := 1; power < num; power *= 10 {

		swapped := (current*10+9)*power + tail
		if swapped > max {
			max = swapped
		}

		tail += remainder * power

		current, remainder = current/10, current%10
	}

	return max
}
