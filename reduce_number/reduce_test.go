package reduce

import "testing"

var examples = []struct {
	want int
	num  int
}{
	{
		want: 6,
		num:  14,
	},
	{
		want: 4,
		num:  8,
	},
	{
		want: 12,
		num:  123,
	},
}

func Test_examples(t *testing.T) {
	for i, e := range examples {
		got := numberOfSteps(e.num)
		if e.want != got {
			t.Fatalf("\ni:%v\nwant:%v\ngot :%v\n", i, e.want, got)
		}
	}
}
