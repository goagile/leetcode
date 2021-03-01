package students

import "testing"

var examples = []struct {
	want      int
	startTime []int
	endTime   []int
	queryTime int
}{
	{
		want:      1,
		startTime: []int{1, 2, 3},
		endTime:   []int{3, 2, 7},
		queryTime: 4,
	},
	{
		want:      1,
		startTime: []int{4},
		endTime:   []int{4},
		queryTime: 4,
	},
	{
		want:      0,
		startTime: []int{4},
		endTime:   []int{4},
		queryTime: 5,
	},
	{
		want:      0,
		startTime: []int{1, 1, 1, 1},
		endTime:   []int{1, 3, 2, 4},
		queryTime: 7,
	},
	{
		want:      5,
		startTime: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		endTime:   []int{10, 10, 10, 10, 10, 10, 10, 10, 10},
		queryTime: 5,
	},
}

func Test_examples_busyStudent(t *testing.T) {
	for i, e := range examples {
		got := busyStudent(e.startTime, e.endTime, e.queryTime)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}
