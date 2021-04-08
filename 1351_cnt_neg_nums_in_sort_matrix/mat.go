package mat

func countNegatives(grid [][]int) int {
	neg := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] < 0 {
				neg++
			}
		}
	}
	return neg
}
