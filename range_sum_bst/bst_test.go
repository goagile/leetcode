package bst

import "testing"

var examples = []struct {
	want int
	low  int
	high int
	root *TreeNode
}{
	{
		want: 32,
		low:  7,
		high: 15,
		root: &TreeNode{Val: 10,
			Left: &TreeNode{Val: 5,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 7},
			},
			Right: &TreeNode{Val: 15,
				Right: &TreeNode{Val: 18},
			},
		},
	},
	{
		want: 23,
		low:  6,
		high: 10,
		root: &TreeNode{Val: 10,
			Left: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 3,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 7,
					Left: &TreeNode{Val: 6},
				},
			},
			Right: &TreeNode{Val: 15,
				Left:  &TreeNode{Val: 13},
				Right: &TreeNode{Val: 18},
			},
		},
	},
}

var bench = examples[0]

func Test_examples_rangeSumBST(t *testing.T) {
	for i, e := range examples {
		got := rangeSumBST(e.root, e.low, e.high)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_rangeSumBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeSumBST(bench.root, bench.low, bench.high)
	}
}

func Test_examples_rangeSumBSTRecursive(t *testing.T) {
	for i, e := range examples {
		got := rangeSumBSTRecursive(e.root, e.low, e.high)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_rangeSumBSTRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeSumBSTRecursive(bench.root, bench.low, bench.high)
	}
}
