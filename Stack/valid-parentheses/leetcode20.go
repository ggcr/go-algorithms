package leetcode20

type stack struct {
	values []rune
}

func (s *stack) push(val rune) {
	s.values = append([]rune{val}, s.values...)
}

func (s *stack) pop() rune {
	res := s.values[0]
	s.values = s.values[1:]
	return res
}

func (s *stack) empty() bool {
	return len(s.values) == 0
}

func (s *stack) get() rune {
	return s.values[0]
}

var paren = map[rune]int32{
	'(': 1,
	')': 2,
	'[': 3,
	']': 4,
	'{': 5,
	'}': 6,
}

func isValid(s string) bool {
	stack := &stack{}
	if len(s) < 2 || len(s)%2 != 0 {
		return false
	}

	for _, v := range s {
		if stack.empty() {
			stack.push(v)
			continue
		}

		// closing parenthesis
		if paren[v]%2 == 0 {
			if paren[stack.get()] != paren[v]-1 {
				return false
			}
			stack.pop()
		} else {
			stack.push(v)
		}
	}
	return stack.empty()
}
