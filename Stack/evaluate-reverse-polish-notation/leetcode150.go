package leetcode150

import (
	"strconv"
)

type Stack struct {
	vals []int
}

var operands = map[string]int{
	"+": 1,
	"-": 2,
	"*": 3,
	"/": 4,
}

func Constructor() Stack {
	return Stack{vals: []int{}}
}

func (s *Stack) Push(val int) {
	s.vals = append([]int{val}, s.vals...)
}

func (s *Stack) Pop() int {
	res := s.vals[0]
	s.vals = s.vals[1:]
	return res
}

func (s *Stack) Top() int {
	return s.vals[0]
}

func evalRPN(tokens []string) int {
	s := Constructor()
	for _, v := range tokens {
		op, ok := operands[v]
		if !ok {
			val, _ := strconv.Atoi(v)
			s.Push(val)
		} else {
			s1 := s.Pop()
			s2 := s.Pop()
			if op == 1 { // addition
				s.Push(s2 + s1)
			} else if op == 2 { // substraction
				s.Push(s2 - s1)
			} else if op == 3 { // multiplication
				s.Push(s2 * s1)
			} else { // division
				s.Push(s2 / s1)
			}
		}
	}
	return s.Top()
}
