package city

import "testing"

var examples = []struct {
	want  string
	paths [][]string
}{
	{
		want: "Sao Paulo",
		paths: [][]string{
			{"London", "New York"},
			{"New York", "Lima"},
			{"Lima", "Sao Paulo"},
		},
	},
	{
		want: "A",
		paths: [][]string{
			{"B", "C"},
			{"D", "B"},
			{"C", "A"},
		},
	},
	{
		want: "Z",
		paths: [][]string{
			{"A", "Z"},
		},
	},
}

func Test_examples_destCity(t *testing.T) {
	for i, e := range examples {
		got := destCity(e.paths)
		if got != e.want {
			t.Fatalf(
				"\ni%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}
