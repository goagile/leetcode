package vowel

import "testing"

var examples = []struct {
	want bool
	s    string
}{
	{want: true, s: "book"},
	{want: false, s: "textbook"},
	{want: false, s: "MerryChristmas"},
	{want: true, s: "AbCdEfGh"},
}

func Test_examples_halvesAreAlike(t *testing.T) {
	for i, e := range examples {
		got := halvesAreAlike(e.s)
		if e.want != got {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\ns:%v\n",
				i, got, e.want, e.s,
			)
		}
	}
}
