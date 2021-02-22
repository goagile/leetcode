package minstack

//
// NewStack
//
func NewStack() *stack {
	return &stack{}
}

type node struct {
	Val  int
	Next *node
}

type stack struct {
	Head  *node
	Count int
}

//
// Push
//
func (s *stack) Push(x int) {
	t := &node{Val: x}
	if s.Head == nil {
		s.Head = t
		s.Count++
		return
	}
	t.Next = s.Head
	s.Head = t
	s.Count++
	return
}

//
// Pop
//
func (s *stack) Pop() (int, bool) {
	if s.Head == nil {
		return 0, false
	}
	t := s.Head
	s.Head = t.Next
	s.Count--
	return t.Val, true
}

//
// Top
//
func (s *stack) Top() (int, bool) {
	if s.Head == nil {
		return 0, false
	}
	return s.Head.Val, true
}

//
// MinStack
//

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		S:   NewStack(),
		Min: NewStack(),
	}
}

type MinStack struct {
	S   *stack
	Min *stack
}

//
// Push
//
func (this *MinStack) Push(x int) {
	if min, ok := this.Min.Top(); !ok {
		this.Min.Push(x)
	} else if x <= min {
		this.Min.Push(x)
	}
	this.S.Push(x)
}

//
// Pop
//
func (this *MinStack) Pop() {
	if a, ok := this.Min.Top(); ok {
		if b, ok := this.S.Top(); ok {
			if a == b {
				this.Min.Pop()
			}
		}
	}
	this.S.Pop()
}

//
// Top
//
func (this *MinStack) Top() int {
	if v, ok := this.S.Top(); ok {
		return v
	}
	return 0
}

//
// GetMin
//
func (this *MinStack) GetMin() int {
	if v, ok := this.Min.Top(); ok {
		return v
	}
	return 0
}
