package xored

import (
	"reflect"
	"testing"
)

var examples = []struct {
	want    []int
	encoded []int
	first   int
}{
	{
		want:    []int{1, 0, 2, 1},
		encoded: []int{1, 2, 3},
		first:   1,
	},
	{
		want:    []int{4, 2, 0, 7, 4},
		encoded: []int{6, 2, 7, 3},
		first:   4,
	},
}

func Test_examples(t *testing.T) {
	for i, e := range examples {
		got := decode(e.encoded, e.first)
		if !reflect.DeepEqual(e.want, got) {
			t.Fatalf("\ni:%v\nwant%v\ngot: %v\n", i, e.want, got)
		}
	}
}
