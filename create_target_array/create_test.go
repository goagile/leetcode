package create

import (
	"reflect"
	"testing"
)

var examples = []struct {
	want  []int
	nums  []int
	index []int
}{
	{
		want:  []int{0, 4, 1, 3, 2},
		nums:  []int{0, 1, 2, 3, 4},
		index: []int{0, 1, 2, 2, 1},
	},
	{
		want:  []int{0, 1, 2, 3, 4},
		nums:  []int{1, 2, 3, 4, 0},
		index: []int{0, 1, 2, 3, 0},
	},
	{
		want:  []int{1},
		nums:  []int{1},
		index: []int{0},
	},
}

var bench = examples[0]

func Test_examples_createTargetArrayList(t *testing.T) {
	for i, e := range examples {
		got := createTargetArrayList(e.nums, e.index)
		if !reflect.DeepEqual(got, e.want) {
			t.Fatalf(
				"\ni:%v\nnums:%v\nindex:%v\nwant:%v\ngot:%v\n",
				i, e.nums, e.index, e.want, got,
			)
		}
	}
}

func Benchmark_createTargetArrayList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createTargetArrayList(bench.nums, bench.index)
	}
}

func Test_examples_createTargetArrayCopy(t *testing.T) {
	for i, e := range examples {
		got := createTargetArrayCopy(e.nums, e.index)
		if !reflect.DeepEqual(got, e.want) {
			t.Fatalf(
				"\ni:%v\nnums:%v\nindex:%v\nwant:%v\ngot:%v\n",
				i, e.nums, e.index, e.want, got,
			)
		}
	}
}

func Benchmark_createTargetArrayCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createTargetArrayCopy(bench.nums, bench.index)
	}
}
