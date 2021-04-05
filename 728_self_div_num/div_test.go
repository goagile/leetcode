package div

import (
	"reflect"
	"testing"
)

var examples = []struct {
	want        []int
	left, right int
}{
	{
		want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		left:  1,
		right: 22,
	},
}

func Test_examples_selfDividingNumbers(t *testing.T) {
	for i, e := range examples {
		got := selfDividingNumbers(e.left, e.right)
		if !reflect.DeepEqual(got, e.want) {
			t.Fatalf(
				"\ni:%v\nwant:%v\ngot:%v\nleft:%v\nright:%v\n",
				i, e.want, got, e.left, e.right,
			)
		}
	}
}
