package max

import "testing"

var examples = []struct {
	want int
	s    string
}{
	{
		want: 3,
		s:    "(1+(2*3)+((8)/4))+1",
	},
	{
		want: 3,
		s:    "(1)+((2))+(((3)))",
	},
	{
		want: 1,
		s:    "1+(2*3)/(2-1)",
	},
	{
		want: 0,
		s:    "1",
	},
}

func Test_examples_maxDepth(t *testing.T) {
	for i, e := range examples {
		got := maxDepth(e.s)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}
