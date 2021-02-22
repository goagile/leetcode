package kids

import (
	"reflect"
	"testing"
)

var examples = []struct {
	want         []bool
	candies      []int
	extraCandies int
}{
	{
		want:         []bool{true, true, true, false, true},
		candies:      []int{2, 3, 5, 1, 3},
		extraCandies: 3,
	},
	{
		want:         []bool{true, false, false, false, false},
		candies:      []int{4, 2, 1, 1, 2},
		extraCandies: 1,
	},
	{
		want:         []bool{true, false, true},
		candies:      []int{12, 1, 12},
		extraCandies: 10,
	},
}

func Test_kidsWithCandies_examples(t *testing.T) {
	for i, e := range examples {
		got := kidsWithCandies(e.candies, e.extraCandies)
		if !reflect.DeepEqual(e.want, got) {
			t.Fatalf(
				"\ni:%v candies:%v extraCandies:%v want:%v got:%v",
				i, e.candies, e.extraCandies, e.want, got,
			)
		}
	}
}
