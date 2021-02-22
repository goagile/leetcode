package ps

import "testing"

var examples = []struct {
	want int
	num  int
}{
	{
		want: 15,
		num:  234,
	},
	{
		want: 21,
		num:  4421,
	},
	{
		want: -12,
		num:  705,
	},
}

func Test_examples(t *testing.T) {
	for i, e := range examples {
		got := subtractProductAndSum(e.num)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}
