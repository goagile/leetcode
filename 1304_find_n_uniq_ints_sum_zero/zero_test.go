package zero

import (
	"reflect"
	"testing"
)

var examples = []struct {
	want []int
	n    int
}{
	{n: 1, want: []int{0}},
	{n: 2, want: []int{-1, 1}},
	{n: 3, want: []int{-1, 0, 1}},
	{n: 4, want: []int{-2, -1, 1, 2}},
}

func Test_examples_sumZero(t *testing.T) {
	for _, e := range examples {
		got := sumZero(e.n)
		if !reflect.DeepEqual(got, e.want) {
			t.Fatalf(
				"\nn:%v\ngot:%v\nwant:%v\n",
				e.n, got, e.want,
			)
		}
	}
}
