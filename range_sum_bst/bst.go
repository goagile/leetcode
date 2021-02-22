package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	s := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		level := []*TreeNode{}
		for _, n := range nodes {
			if low <= n.Val && n.Val <= high {
				s += n.Val
			}
			if n.Left != nil {
				level = append(level, n.Left)
			}
			if n.Right != nil {
				level = append(level, n.Right)
			}
		}
		nodes = level
	}
	return s
}

func rangeSumBSTRecursive(n *TreeNode, low int, high int) int {
	s := 0
	if n == nil {
		return s
	}
	if low <= n.Val && n.Val <= high {
		s += n.Val
	}
	s += rangeSumBSTRecursive(n.Left, low, high)
	s += rangeSumBSTRecursive(n.Right, low, high)
	return s
}
