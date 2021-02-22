package matches

import "testing"

var examples = []struct {
	want int
	n    int
}{
	{
		want: 6,
		n:    7,
	},
	{
		want: 13,
		n:    14,
	},
}

func Test_examples_numberOfMatches(t *testing.T) {
	for i, e := range examples {
		got := numberOfMatches(e.n)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Test_examples_numberOfMatchesR(t *testing.T) {
	for i, e := range examples {
		got := numberOfMatchesR(e.n)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}
