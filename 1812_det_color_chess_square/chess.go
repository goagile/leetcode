package chess

func squareIsWhite(c string) bool {
	x, y := c[0], c[1]
	return (x%2 == 0 && y%2 != 0) ||
		(x%2 != 0 && y%2 == 0)
}
