package numrect

func countGoodRectangles(rectangles [][]int) int {
	good := 0
	max := 0
	for _, r := range rectangles {
		side := r[0]
		if side > r[1] {
			side = r[1]
		}
		if side > max {
			max = side
			good = 1
		} else if side == max {
			good++
		}
	}
	return good
}
