package match

import "testing"

var examples = []struct {
	want      int
	items     [][]string
	ruleKey   string
	ruleValue string
}{
	{
		want: 1,
		items: [][]string{
			{"phone", "blue", "pixel"},
			{"computer", "silver", "lenovo"},
			{"phone", "gold", "iphone"},
		},
		ruleKey: "color", ruleValue: "silver",
	},
	{
		want: 2,
		items: [][]string{
			{"phone", "blue", "pixel"},
			{"computer", "silver", "lenovo"},
			{"phone", "gold", "iphone"},
		},
		ruleKey: "type", ruleValue: "phone",
	},
}

var bench = examples[0]

func Test_examples_countMatches(t *testing.T) {
	for i, e := range examples {
		got := countMatches(e.items, e.ruleKey, e.ruleValue)
		if got != e.want {
			t.Fatalf(
				"\ni%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_countMatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countMatches(bench.items, bench.ruleKey, bench.ruleValue)
	}
}

func Test_examples_countMatchesHash(t *testing.T) {
	for i, e := range examples {
		got := countMatchesHash(e.items, e.ruleKey, e.ruleValue)
		if got != e.want {
			t.Fatalf(
				"\ni%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_countMatchesHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countMatchesHash(bench.items, bench.ruleKey, bench.ruleValue)
	}
}
