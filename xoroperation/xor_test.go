package xor

import "testing"

var examples = []struct {
	want  int
	n     int
	start int
}{
	{
		want:  8,
		n:     5,
		start: 0,
	},
	{
		want:  8,
		n:     4,
		start: 3,
	},
	{
		want:  7,
		n:     1,
		start: 7,
	},
	{
		want:  2,
		n:     10,
		start: 5,
	},
}

func Test_examples_xorOperation(t *testing.T) {
	for i, e := range examples {
		got := xorOperation(e.n, e.start)
		if got != e.want {
			t.Fatalf("\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}
