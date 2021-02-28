package matrix

func diagonalSum(A [][]int) int {
	s, n := 0, len(A)
	for i := 0; i < n; i++ {
		s += A[i][i]
		j := n - i - 1
		if i != j {
			s += A[i][j]
		}
	}
	return s
}

func diagonalSumOdd(A [][]int) int {
	s, n := 0, len(A)
	for i := 0; i < n; i++ {
		s += A[i][i] + A[i][n-i-1]
	}
	if n%2 == 1 {
		s -= A[n/2][n/2]
	}
	return s
}
