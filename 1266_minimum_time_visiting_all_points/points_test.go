package points

import (
	"testing"
)

var examples = []struct {
	want   int
	points [][]int
}{
	{
		want: 7,
		points: [][]int{
			{1, 1}, {3, 4}, {-1, 0},
		},
	},
	{
		want: 5,
		points: [][]int{
			{3, 2}, {-2, 2},
		},
	},
	{
		want: 0,
		points: [][]int{
			{0, 0}, {0, 0},
		},
	},
	{
		want: 4,
		points: [][]int{
			{0, 0}, {1, 0}, {1, 1}, {0, 1}, {0, 0},
		},
	},
	{
		want: 1,
		points: [][]int{
			{0, 0}, {-1, -1},
		},
	},
}

func Test_examples_minTimeToVisitAllPoints(t *testing.T) {
	for i, e := range examples {
		got := minTimeToVisitAllPoints(e.points)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\npoints:%v\n",
				i, got, e.want, e.points,
			)
		}
	}
}
