package city

func destCity(paths [][]string) string {
	cities := make(map[string]string)
	for _, path := range paths {
		cities[path[0]] = path[1]
	}
	for _, dest := range cities {
		if _, found := cities[dest]; !found {
			return dest
		}
	}
	return ""
}
