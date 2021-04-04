package flip

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)
	for _, a := range image {
		for i, j := 0, n-1; i <= n/2 && j >= n/2; i, j = i+1, j-1 {
			if i != j {
				a[i], a[j] = a[j]^1, a[i]^1
			} else {
				a[i] ^= 1
			}
		}
	}
	return image
}
