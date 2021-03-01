package numrect

import "testing"

var examples = []struct {
	want       int
	rectangles [][]int
}{
	{
		want: 3,
		rectangles: [][]int{
			{5, 8}, {3, 9}, {5, 12}, {16, 5},
		},
	},
	{
		want: 3,
		rectangles: [][]int{
			{2, 3}, {3, 7}, {4, 3}, {3, 7},
		},
	},
}

func Test_examples_countGoodRectangles(t *testing.T) {
	for i, e := range examples {
		got := countGoodRectangles(e.rectangles)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}
