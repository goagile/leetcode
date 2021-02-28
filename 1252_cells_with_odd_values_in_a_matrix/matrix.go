package matrix

func oddCells(n int, m int, indices [][]int) int {
	A := make([][]int, n, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m, m)
	}
	odds := 0
	for _, nd := range indices {
		for c := 0; c < m; c++ {
			r := nd[0]
			A[r][c]++
			if A[r][c]%2 != 0 {
				odds++
			} else {
				odds--
			}
		}
		for r := 0; r < n; r++ {
			c := nd[1]
			A[r][c]++
			if A[r][c]%2 != 0 {
				odds++
			} else {
				odds--
			}
		}
	}
	return odds
}
