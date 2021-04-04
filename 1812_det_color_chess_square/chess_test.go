package chess

import "testing"

var examples = []struct {
	want bool
	c    string
}{
	{want: false, c: "a1"},
	{want: true, c: "h3"},
	{want: false, c: "c7"},
}

func Test_examples_squareIsWhite(t *testing.T) {
	for i, e := range examples {
		got := squareIsWhite(e.c)
		if e.want != got {
			t.Fatalf(
				"\ni:%v\nwant:%v\ngot:%v\n",
				i, e.want, got,
			)
		}
	}
}
