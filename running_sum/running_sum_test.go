package runningsum

import (
	"reflect"
	"testing"
)

var examples = []struct {
	arr  []int
	want []int
}{
	{
		arr:  []int{5, 5, 5},
		want: []int{5, 10, 15},
	},
	{
		arr:  []int{1, 2, 3, 4},
		want: []int{1, 3, 6, 10},
	},
	{
		arr:  []int{1, 1, 1, 1, 1},
		want: []int{1, 2, 3, 4, 5},
	},
}

func Test_runningSum(t *testing.T) {
	for i, e := range examples {
		got := runningSum(e.arr)
		if !reflect.DeepEqual(e.want, got) {
			t.Fatalf("\ni:%v\nwant:%v\ngot: %v", i, e.want, got)
		}
	}
}

func Benchmark_runningSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runningSum([]int{1, 2, 3, 4})
	}
}

func Test_runningSumShort(t *testing.T) {
	for i, e := range examples {
		got := runningSumShort(e.arr)
		if !reflect.DeepEqual(e.want, got) {
			t.Fatalf("\ni:%v\nwant:%v\ngot: %v", i, e.want, got)
		}
	}
}

func Benchmark_runningSumShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runningSumShort([]int{1, 2, 3, 4})
	}
}
