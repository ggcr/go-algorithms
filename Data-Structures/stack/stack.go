package stack

type stack struct {
	values []int
}

func (s *stack) push(val int) {
	s.values = append([]int{val}, s.values...)
}

func (s *stack) pop() int {
	res := s.values[0]
	s.values = s.values[1:]
	return res
}

func (s *stack) empty() bool {
	return len(s.values) == 0
}
