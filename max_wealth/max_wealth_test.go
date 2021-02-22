package maxwealth

import (
	"testing"
)

var examples = []struct {
	want     int
	accounts [][]int
}{
	{
		want: 6,
		accounts: [][]int{
			{1, 2, 3},
			{3, 2, 1},
		},
	},
	{
		want: 10,
		accounts: [][]int{
			{1, 5},
			{7, 3},
			{3, 5},
		},
	},
	{
		want: 17,
		accounts: [][]int{
			{2, 8, 7},
			{7, 1, 3},
			{1, 9, 5},
		},
	},
}

func Benchmark_maximumWealthRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumWealthRange(examples[0].accounts)
	}
}

func Test_maximumWealthRange(t *testing.T) {
	for i, e := range examples {
		got := maximumWealthRange(e.accounts)
		if e.want != got {
			t.Fatalf("\ni:%v want:%v got:%v\n", i, e.want, got)
		}
	}
}

func Benchmark_maximumWealthLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumWealthLoop(examples[0].accounts)
	}
}

func Test_maximumWealthLoop(t *testing.T) {
	for i, e := range examples {
		got := maximumWealthLoop(e.accounts)
		if e.want != got {
			t.Fatalf("\ni:%v want:%v got:%v\n", i, e.want, got)
		}
	}
}
