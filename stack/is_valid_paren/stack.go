package stack

//
// isValid
//
func isValid(s string) bool {
    r := NewStack()
    for _, t := range s {
    	switch t {
    	
    	case '{', '[', '(':
    		r.Push(t)

    	case '}':
    		if v := r.Pop(); v != '{' {
    			return false
    		}

    	case ']':
    		if v := r.Pop(); v != '[' {
    			return false
    		}

    	case ')':
    		if v := r.Pop(); v != '(' {
    			return false
    		}

    	default:
    		continue
    	}
    }
    if !r.Empty() {
    	return false
    }
    return true
}

//
// Stack
//
func NewStack() *stack {
	return &stack{}
}

type stack struct {
	head *node
	count int
}

type node struct {
	val rune
	next *node
}

func (s *stack) Push(v rune) {
	n := &node{val: v}
	if s.head == nil {
		s.head = n
		s.count++
		return
	}
	t := s.head
	s.head = n
	n.next = t
	s.count++
	return
} 

func (s *stack) Empty() bool {
	return 0 == s.count
}

func (s *stack) Pop() rune {
	if s.Empty() {
		return '0'
	}
	t := s.head
	s.head = t.next
	s.count--
	return t.val
}
