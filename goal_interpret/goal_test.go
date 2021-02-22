package goal

import "testing"

var examples = []struct {
	want    string
	command string
}{
	{
		want:    "Goal",
		command: "G()(al)",
	},
	{
		want:    "Gooooal",
		command: "G()()()()(al)",
	},
	{
		want:    "alGalooG",
		command: "(al)G(al)()()G",
	},
}

var bench = examples[0]

func Test_examples_interpret(t *testing.T) {
	for i, e := range examples {
		got := interpret(e.command)
		if e.want != got {
			t.Fatalf(
				"\ni:%v\ncommand:%v\nwant:%v\ngot :%v\n",
				i, e.command, e.want, got,
			)
		}
	}
}

func Benchmark_interpret(b *testing.B) {
	for i := 0; i < b.N; i++ {
		interpret(bench.command)
	}
}

func Test_examples_interpretBuiltin(t *testing.T) {
	for i, e := range examples {
		got := interpretBuiltin(e.command)
		if e.want != got {
			t.Fatalf(
				"\ni:%v\ncommand:%v\nwant:%v\ngot :%v\n",
				i, e.command, e.want, got,
			)
		}
	}
}

func Benchmark_interpretBuiltin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		interpretBuiltin(bench.command)
	}
}

func Test_examples_interpretMap(t *testing.T) {
	for i, e := range examples {
		got := interpretMap(e.command)
		if e.want != got {
			t.Fatalf(
				"\ni:%v\ncommand:%v\nwant:%v\ngot :%v\n",
				i, e.command, e.want, got,
			)
		}
	}
}

func Benchmark_interpretMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		interpretMap(bench.command)
	}
}
