package flip

import (
	"reflect"
	"testing"
)

var examples = []struct {
	want  [][]int
	image [][]int
}{
	{
		want: [][]int{
			{1, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		},
		image: [][]int{
			{1, 1, 0},
			{1, 0, 1},
			{0, 0, 0},
		},
	},
	{
		want: [][]int{
			{1, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 1},
			{1, 0, 1, 0},
		},
		image: [][]int{
			{1, 1, 0, 0},
			{1, 0, 0, 1},
			{0, 1, 1, 1},
			{1, 0, 1, 0},
		},
	},
}

func Test_flipAndInvertImage_examples(t *testing.T) {
	for i, e := range examples {
		got := flipAndInvertImage(e.image)
		if !reflect.DeepEqual(e.want, got) {
			t.Fatalf(
				"\ni:%v\nwant:%v\ngot:%v\n",
				i, e.want, got,
			)
		}
	}
}
