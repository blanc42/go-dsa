package stack

type Stack struct {
	data []int
}

func (s *Stack) Push(d int) {
	s.data = append(s.data, d)
}

func (s *Stack) Pop() int {
	if len(s.data) != 0 {
		out := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return out
	}
	return 0
}

// What should the 'pop()' method return when the stack is empty?
// https://stackoverflow.com/a/7390181
func (s *Stack) IsEmpty() bool {
	if len(s.data) == 0 {
		return true
	} else {
		return false
	}
}
