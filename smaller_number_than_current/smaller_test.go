package smaller

import (
	"reflect"
	"testing"
)

var examples = []struct {
	want []int
	nums []int
}{
	{
		want: []int{4, 0, 1, 1, 3},
		nums: []int{8, 1, 2, 2, 3},
	},
	{
		want: []int{2, 1, 0, 3},
		nums: []int{6, 5, 4, 8},
	},
	{
		want: []int{0, 0, 0, 0},
		nums: []int{7, 7, 7, 7},
	},
}

func Test_examples(t *testing.T) {
	for i, e := range examples {
		got := smallerNumbersThanCurrent(e.nums)
		if !reflect.DeepEqual(e.want, got) {
			t.Fatalf("\ni:%v\nwant:%v\ngot :%v\n", i, e.want, got)
		}
	}
}
