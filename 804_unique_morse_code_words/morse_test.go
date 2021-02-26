package morse

import "testing"

var examples = []struct {
	want  int
	words []string
}{
	{
		want:  2,
		words: []string{"gin", "zen", "gig", "msg"},
	},
}

func Test_examples_uniqueMorseRepresentations(t *testing.T) {
	for i, e := range examples {
		got := uniqueMorseRepresentations(e.words)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}
