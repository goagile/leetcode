package maximum

import "testing"

var examples = []struct {
	want int
	num  int
}{
	{
		want: 9999,
		num:  9996,
	},
	{
		want: 9996,
		num:  9966,
	},
	{
		want: 9966,
		num:  9666,
	},
	{
		want: 966,
		num:  666,
	},
	{
		want: 99,
		num:  69,
	},
	{
		want: 9,
		num:  6,
	},
}

func Test_examples_maximum69Number(t *testing.T) {
	for i, e := range examples {
		got := maximum69Number(e.num)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}
