package linklist

import (
	"fmt"
	"strings"
)

//
// ListNode
//
type ListNode struct {
	Val int
	Next *ListNode
}

func (n *ListNode) String() string {
	return fmt.Sprintf("%v", n.Val)
}

func (n *ListNode) Eq(other *ListNode) bool {
	if other == nil {
		return false
	}
	return (n.Val == other.Val) 
}

//
// List
//
type List struct {
	Head *ListNode
	Count int
}

func New() *List {
	return &List{}
}

func NewFromNode(n *ListNode) *List {
	l := New()
	for n != nil {
		l.Append(n.Val)
		n = n.Next
	}
	return l
}

func NewFromElems(elms ...int) *List {
	l := New()
	for _, e := range elms {
		l.Append(e)
	}
	return l
}

func (l *List) Get(v int) *ListNode {
	n := l.Head
	for n != nil {
		if v == n.Val {
			return n
		}
		n = n.Next
	}
	return nil
}

func (l *List) String() string {
	var s []string
	n := l.Head
	for n != nil {
		s = append(s, fmt.Sprintf("%v", n.Val))
		n = n.Next
	}
	return strings.Join(s, "->")
}

func (l *List) Eq(other *List) bool {
	if l.Count != other.Count {
		return false
	}
	a := l.Head
	b := other.Head
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return true
}

func (l *List) Append(v int) *ListNode {
	t := &ListNode{Val: v}
	if l.Empty() {
		l.Head = &ListNode{Val: v}
		l.Count++
		return l.Head
	}
	n := l.Head
	var p *ListNode
	for n != nil {
		p = n
 		n = n.Next
	}
	p.Next = t
	l.Count++
	return t
}

func (l *List) Empty() bool {
	return 0 == l.Count
}

func (l *List) RemoveDuplicates() *List {
	return NewFromNode(deleteDuplicates(l.Head))
}

func (l *List) Reversed() *List {
	return NewFromNode(reverseList(l.Head))
}

func (l *List) Copy() *List {
	return NewFromNode(copy(l.Head))
}

func (l *List) IsPalindrome() bool {
	return isPalindrome(l.Head)
}

func (l *List) Decimal() int {
	return getDecimalValue(l.Head)
}

//
// delete Dulicates
//
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	n := head
	for n != nil {
		if n.Next != nil && n.Next.Val == n.Val {
			n.Next = n.Next.Next
			continue
		}
		n = n.Next
	}
	return head
}

//
// reverse list
//
func reverseList(head *ListNode) *ListNode {
    c := head
    var n, p *ListNode
    for c != nil {
    	n = c.Next
    	c.Next = p
    	p = c
    	c = n
    }
    return p
}

func copy(head *ListNode) *ListNode {
	n := head
	copyhead := &ListNode{Val: n.Val}
	c := copyhead
	for n.Next != nil {
		c.Next = &ListNode{Val: n.Next.Val}
		n = n.Next
		c = c.Next
	}
	return copyhead
}

//
// palindrome
//
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	c := copy(head)
	r := reverseList(c)
	n := head
	for n != nil {
		if 0 != (n.Val-r.Val) {
			return false
		}
		n = n.Next
		r = r.Next
	}
	return true
}

func getDecimalValue(head *ListNode) int {
	v := 0
	for head != nil {
		v = 2 * v + head.Val
		head = head.Next
	}
	return v
}

func deleteNode(node *ListNode) {
    n := node.Next
    nn := node.Next.Next
    node.Val = n.Val
    node.Next = nn
    n.Next = nil
}

func removeElements(head *ListNode, val int) *ListNode {
    n := head
    if n == nil {
    	return nil
    }
    var p *ListNode
    for n != nil && n.Next != nil {
    	if n.Next != nil && val == n.Val {
    		n.Val = n.Next.Val
    		n.Next = n.Next.Next
    		continue
    	}
    	p = n
    	n = n.Next
    }
    if p != nil && p.Next != nil && p.Next.Val == val {
    	p.Next = nil
    }
    if head.Val == val {
    	return nil
    }
    return head
}

func middleNode(head *ListNode) *ListNode {
    n := head
    if n == nil {
    	return nil
    }
    if n.Next == nil {
    	return n
    }
    var midhead, t *ListNode
    i := 0
    c := count(n)
    m := int(c/2)
    for n != nil {
    	if i == m {
			midhead	= &ListNode{Val:n.Val}
			t = midhead
    	} else if i > m {
    		t.Next = &ListNode{Val:n.Val}
    		t = t.Next
    	}
    	i++
    	n = n.Next
    }
    return midhead
}

//
// cmp count
// 
func cmpcount(a, b *ListNode) int {
	ha, hb := a, b
	diff := count(ha) - count(hb) 
	return diff
}

//
// count
// 
func count(head *ListNode) int {
	c := 0
	n := head
	for n != nil {
		c++
		n = n.Next
	}
	return c
}

//
// scroll
//
func scroll(a *ListNode, n int) *ListNode {
	if n < 0 {
		n = -n
	}

	for i := 0; a != nil && i < n; i++ {
		a = a.Next
	}

	return a
}

//
// has cycle
//
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	c := head
	n := c
	var p *ListNode
	for c.Next != nil {
		n = c.Next
		c.Next = p
		p = c
		c = n
	}
	return (c == head)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	if a == nil || b == nil {
		return nil
	}

	if diff := cmpcount(a, b); diff > 0 {
		a = scroll(a, diff)
	} else if diff < 0 {
		b = scroll(b, diff)
	}

	if a == b {
		return a
	}

	for a != nil && b != nil {
		if a.Next == b.Next {
			return a.Next
		}
		a, b = a.Next, b.Next
	}

	return nil
}

//
// sortedadd
//
func sortedadd(head *ListNode, v int) *ListNode {
	t := &ListNode{Val: v}
	if head == nil {
		return t
	}
	if v < head.Val {
		t.Next = head
		head = t
		return head
	}
	n := head
	var p *ListNode
	for n != nil && v >= n.Val {
		p = n
		n = n.Next
	}
	t.Next = p.Next
	p.Next = t
	return head
}

//
// merge
// 
func merge(headA, headB *ListNode) *ListNode {
	var s *ListNode

	a := headA
	b := headB

	for {
		if a != nil {
			s = sortedadd(s, a.Val)
			a = a.Next
		}
		if b != nil {
			s = sortedadd(s, b.Val)
			b = b.Next
		}
		if a == nil && b == nil {
			break
		}
	}

	return s
}
