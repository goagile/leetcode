package alph

import "testing"

var examples = []struct {
	want string
	s    string
}{
	{
		want: "jkab",
		s:    "10#11#12",
	},
	{
		want: "acz",
		s:    "1326#",
	},
	{
		want: "y",
		s:    "25#",
	},
	{
		want: "abcdefghijklmnopqrstuvwxyz",
		s:    "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#",
	},
}

func Test_examples_freqAlphabets(t *testing.T) {
	for i, e := range examples {
		got := freqAlphabets(e.s)
		if e.want != got {
			t.Fatalf(
				"\ni:%v\nwant:%v\ngot:%v\ns:%v\n",
				i, e.want, got, e.s,
			)
		}
	}
}
