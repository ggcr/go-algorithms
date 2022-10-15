package leetcode22

type stack struct {
	values []rune
}

func (s *stack) pushLast(val rune) {
	s.values = append(s.values, val)
}

func (s *stack) popLast() rune {
	res := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return res
}

var res = []string{}

func backtracking(openP, closeP, n int, s *stack) []string {
	if openP == closeP && openP == n {
		res = append(res, string(s.values))
		return res
	} else {

		if openP < n {
			s.pushLast('(')
			backtracking(openP+1, closeP, n, s)
			s.popLast()
		}

		if closeP < openP {
			s.pushLast(')')
			backtracking(openP, closeP+1, n, s)
			s.popLast()
		}
	}

	return nil
}

func generateParenthesis(n int) []string {
	res = []string{}
	openP := 0
	closeP := 0
	s := &stack{}
	backtracking(openP, closeP, n, s)
	return res
}
