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

func countNegativesBisect(grid [][]int) int {
	c := 0
	for _, r := range grid {
		c += bisect(r)
	}
	return c
}

func bisect(a []int) int {
	n := len(a)
	i := n
	l, h := 0, (n - 1)
	for l <= h {
		m := (l + h) / 2
		if a[m] >= 0 {
			l = m + 1
		} else if a[m] < 0 {
			h = m - 1
			i = m
		}
	}
	return n - i
}
