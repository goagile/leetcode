package trunc

import "testing"

var examples = []struct {
	want string
	s    string
	k    int
}{
	{
		want: "Hello how are you",
		s:    "Hello how are you Contestant",
		k:    4,
	},
	{
		want: "What is the solution",
		s:    "What is the solution to this problem",
		k:    4,
	},
	{
		want: "chopper is not a tanuki",
		s:    "chopper is not a tanuki",
		k:    5,
	},
}

func Test_examples_truncateSentence(t *testing.T) {
	for i, e := range examples {
		got := truncateSentence(e.s, e.k)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\nwant:%v\ngot:%v\ns:%v\n",
				i, e.want, got, e.s,
			)
		}
	}
}
