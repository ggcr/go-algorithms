package leetcode853

import (
	"sort"
)

type stack struct {
	values []float64
}

func Constructor(vector []float64) stack {
	return stack{values: vector}
}

func (s *stack) push(val float64) {
	s.values = append([]float64{val}, s.values...)
}

func (s *stack) pushLast(val float64) {
	s.values = append(s.values, val)
}

func (s *stack) pop() {
	s.values = s.values[1:]
}

func (s *stack) top() float64 {
	return s.values[0]
}

func (s *stack) empty() bool {
	return len(s.values) == 0
}

func GetMapSpeed(position, speed []int) map[int]int {
	res := make(map[int]int, len(position))
	for k, v := range position {
		res[v] = speed[k]
	}
	return res
}

func carFleet(target int, position []int, speed []int) int {
	speedMap := GetMapSpeed(position, speed)
	sort.Slice(position, func(i, j int) bool {
		return position[i] > position[j]
	})
	aux := Constructor([]float64{})
	res := 1
	for i := 0; i < len(position); i++ {
		// Get num of iterations to get to the target
		n_it := float64(target-position[i]) / float64(speedMap[position[i]])
		if !aux.empty() && n_it > aux.top() {
			// Pop
			res += 1
			for len(aux.values) > 0 && n_it > aux.top() {
				aux.pop()
			}
		}
		aux.pushLast(n_it)
	}
	return res
}
