package sum

func sumOddLengthSubarrays(a []int) int {
	s := 0
	for k := 1; k <= len(a); k += 2 {
		for i := 0; (i + k) <= len(a); i++ {
			for _, v := range a[i : i+k] {
				s += v
			}
		}
	}
	return s
}

func sumOddLengthSubarraysFun(arr []int) int {
	s := 0
	for oddLen := 1; oddLen <= len(arr); oddLen += 2 {
		s += sumOfAllSubarrays(arr, oddLen)
	}
	return s
}

func sumOfAllSubarrays(arr []int, oddLen int) int {
	s := 0
	for i := 0; (i + oddLen) <= len(arr); i++ {
		subArray := arr[i:(i + oddLen)]
		s += sumOfElements(subArray)
	}
	return s
}

func sumOfElements(subArr []int) int {
	s := 0
	for _, elem := range subArr {
		s += elem
	}
	return s
}
