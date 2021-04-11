package discount

import (
	"reflect"
	"testing"
)

var examples = []struct {
	want   []int
	prices []int
}{
	{
		want:   []int{4, 2, 4, 2, 3},
		prices: []int{8, 4, 6, 2, 3},
	},
	{
		want:   []int{1, 2, 3, 4, 5},
		prices: []int{1, 2, 3, 4, 5},
	},
	{
		want:   []int{9, 0, 1, 6},
		prices: []int{10, 1, 1, 6},
	},
}

func Test_examples_finalPrices(t *testing.T) {
	for i, e := range examples {
		got := finalPrices(e.prices)
		if !reflect.DeepEqual(got, e.want) {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\nprices:%v\n",
				i, got, e.want, e.prices,
			)
		}
	}
}
