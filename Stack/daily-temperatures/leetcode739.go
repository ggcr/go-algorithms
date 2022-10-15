package leetcode739

type stack struct {
	values []int
}

func (s *stack) push(val int) {
	s.values = append([]int{val}, s.values...)
}

func (s *stack) pop() {
	s.values = s.values[1:]
}

func (s *stack) top() int {
	return s.values[0]
}

func (s *stack) empty() bool {
	return len(s.values) == 0
}

func dailyTemperatures(temperatures []int) []int {
	res := &stack{}
	idx := &stack{}

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(idx.values) > 0 && temperatures[idx.top()] <= temperatures[i] {
			idx.pop()
		}

		if idx.empty() {
			res.push(0)
		} else {
			res.push(idx.top() - i)
		}

		idx.push(i)
	}

	return res.values
}
