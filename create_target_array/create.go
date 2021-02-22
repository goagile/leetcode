package create

func createTargetArrayCopy(nums []int, index []int) []int {
	n := len(nums)
	z := make([]int, n, n)
	for i := 0; i < n; i++ {
		k, v := index[i], nums[i]
		copy(z[k+1:], z[k:])
		z[k] = v
	}
	return z
}

func createTargetArrayList(nums []int, index []int) []int {
	l := &List{}
	for i := 0; i < len(nums); i++ {
		l.Insert(nums[i], index[i])
	}
	z := make([]int, 0)
	for n := l.Head; n != nil; n = n.Next {
		z = append(z, n.Val)
	}
	return z
}

type List struct {
	Head  *node
	Count int
}

type node struct {
	Next *node
	Val  int
}

func (l *List) Insert(v int, k int) {
	m := &node{Val: v}
	n := l.Head
	if k == 0 {
		m.Next = n
		l.Head = m
		l.Count++
		return
	}
	r := (k - l.Count)
	if r > 0 {
		if n == nil {
			l.Head = &node{}
			l.Count++
			n = l.Head
			r--
		}
		for i := 0; i < r; i++ {
			n.Next = &node{}
			n = n.Next
			l.Count++
		}
	} else {
		for i := 0; i < k-1; i++ {
			n = n.Next
		}
	}
	if n.Next != nil {
		m.Next = n.Next
	}
	n.Next = m
	l.Count++
}
