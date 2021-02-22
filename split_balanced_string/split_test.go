package split

import "testing"

var examples = []struct {
	want int
	s    string
}{
	{
		want: 4,
		s:    "RLRRLLRLRL",
	},
	{
		want: 3,
		s:    "RLLLLRRRLR",
	},
	{
		want: 1,
		s:    "LLLLRRRR",
	},
	{
		want: 2,
		s:    "RLRRRLLRLL",
	},
}

func Test_examples_balancedStringSplit(t *testing.T) {
	for i, e := range examples {
		got := balancedStringSplit(e.s)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\ns:%v\n",
				i, got, e.want, e.s,
			)
		}
	}
}
