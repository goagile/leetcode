package match

var fields = map[string]int{
	"type":  0,
	"color": 1,
	"name":  2,
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	c := 0
	for _, item := range items {
		if ruleValue == item[fields[ruleKey]] {
			c++
		}
	}
	return c
}

func countMatchesHash(items [][]string, ruleKey string, ruleValue string) int {
	var index = map[string]map[string]int{
		"type":  {},
		"color": {},
		"name":  {},
	}
	for _, item := range items {
		index["type"][item[0]]++
		index["color"][item[1]]++
		index["name"][item[2]]++
	}
	return index[ruleKey][ruleValue]
}
