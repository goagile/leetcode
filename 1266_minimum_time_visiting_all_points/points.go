package points

func minTimeToVisitAllPoints(ps [][]int) int {
	var t int
	for i := 1; i < len(ps); i++ {
		t += move(ps[i-1], ps[i])
	}
	return t
}

const (
	x, y = 0, 1
)

func move(a, b []int) int {
	var n int
	tx, nx := a[x], abs(b[x]-a[x])
	ty, ny := a[y], abs(b[y]-a[y])
	for ; n < nx || n < ny; n++ {
		tx += step(tx, b[x])
		ty += step(ty, b[y])
	}
	return n
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func step(a, b int) int {
	r := b - a
	if r > 0 {
		return 1
	} else if r < 0 {
		return -1
	}
	return 0
}
