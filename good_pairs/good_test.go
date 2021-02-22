package good

import "testing"

var examples = []struct {
	want int
	arr  []int
}{
	{
		want: 4,
		arr:  []int{1, 2, 3, 1, 1, 3},
	},
	{
		want: 6,
		arr:  []int{1, 1, 1, 1},
	},
	{
		want: 0,
		arr:  []int{1, 2, 3},
	},
}

func Test_numIdenticalPairs(t *testing.T) {
	for i, e := range examples {
		got := numIdenticalPairs(e.arr)
		if e.want != got {
			t.Fatalf("\ni:%v want:%v got:%v", i, e.want, got)
		}
	}
}

func Benchmark_numIdenticalPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numIdenticalPairs(examples[0].arr)
	}
}

func Test_numIdenticalPairsHash(t *testing.T) {
	for i, e := range examples {
		got := numIdenticalPairsHash(e.arr)
		if e.want != got {
			t.Fatalf("\ni:%v want:%v got:%v", i, e.want, got)
		}
	}
}

func Benchmark_numIdenticalPairsHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numIdenticalPairsHash(examples[0].arr)
	}
}
