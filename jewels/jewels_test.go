package jewels

import "testing"

var examples = []struct {
	want   int
	jewels string
	stones string
}{
	{want: 3, jewels: "aA", stones: "aAAbbbb"},
	{want: 0, jewels: "z", stones: "ZZZ"},
	{
		want:   12,
		jewels: "aA",
		stones: `aAAbbbbaAAbbbbaAAbbbbaAAbbbb`,
	},
}

var bench = examples[2]

//
// Nested Loops
//
func Test_numJewelsInStones(t *testing.T) {
	for i, e := range examples {
		got := numJewelsInStones(e.jewels, e.stones)
		if e.want != got {
			t.Fatalf(
				"\ni:%v got:%v want:%v jewels:%v stones:%v",
				i, got, e.want, e.jewels, e.stones)
		}
	}
}

func Benchmark_numJewelsInStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numJewelsInStones(bench.jewels, bench.stones)
	}
}

//
// Contains
//

func Test_numJewelsInStonesContains(t *testing.T) {
	for i, e := range examples {
		got := numJewelsInStonesContains(e.jewels, e.stones)
		if e.want != got {
			t.Fatalf(
				"\ni:%v got:%v want:%v jewels:%v stones:%v",
				i, got, e.want, e.jewels, e.stones)
		}
	}
}

func Benchmark_numJewelsInStonesContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numJewelsInStonesContains(bench.jewels, bench.stones)
	}
}

//
// Hash Map
//
func Test_numJewelsInStonesHashMap(t *testing.T) {
	for i, e := range examples {
		got := numJewelsInStonesHashMap(e.jewels, e.stones)
		if e.want != got {
			t.Fatalf(
				"\ni:%v got:%v want:%v jewels:%v stones:%v",
				i, got, e.want, e.jewels, e.stones)
		}
	}
}

func Benchmark_numJewelsInStonesHashMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numJewelsInStonesHashMap(bench.jewels, bench.stones)
	}
}
