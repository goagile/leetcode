package stack

//
// minOperations
//
func minOperations(logs []string) int {

	s := NewStack()

	for _, x := range logs {
		switch x {
		case "./":
			continue
		case "../":
			s.Pop()
		default:
			s.Push("../")
		}
	}

	return s.Count()

}

//
// Count Stack
//
func NewStack() *stack {
	return &stack{}
}

type stack struct {
	count int
}

func (s *stack) Push(x string) {
	s.count++
}

func (s *stack) Pop() {
	if s.count > 0 {
		s.count--
	}
}

func (s *stack) Count() int {
	return s.count
}
