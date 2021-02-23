package tolowcase

import "testing"

var examples = []struct {
	want string
	str  string
}{
	{
		want: "waaaagh#$@$@>iiiibnooz",
		str:  "WaaAAGh#$@$@>iIiIbNooZ",
	},
	{
		want: "lovely",
		str:  "LOVELY",
	},
	{
		want: "hello",
		str:  "Hello",
	},
	{
		want: "here",
		str:  "here",
	},
}

var bench = examples[0]

func Test_examples_toLowerCaseStringsBuilder(t *testing.T) {
	for i, e := range examples {
		got := toLowerCaseStringsBuilder(e.str)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_toLowerCaseStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toLowerCaseStringsBuilder(bench.str)
	}
}

func Test_examples_toLowerCaseStd(t *testing.T) {
	for i, e := range examples {
		got := toLowerCaseStd(e.str)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_toLowerCaseStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toLowerCaseStd(bench.str)
	}
}

func Test_examples_toLowerCaseBytesBuffer(t *testing.T) {
	for i, e := range examples {
		got := toLowerCaseBytesBuffer(e.str)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_toLowerCaseBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toLowerCaseBytesBuffer(bench.str)
	}
}
