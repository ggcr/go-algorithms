package leetcode155

import "math"

type MinStack struct {
	vals []int
	min  []int
}

var minElement = math.MaxInt

func Constructor() MinStack {
	minElement = math.MaxInt
	return MinStack{vals: []int{}, min: []int{}}
}

func (this *MinStack) Push(val int) {
	this.vals = append([]int{val}, this.vals...)

	if val < minElement {
		minElement = val
	}
	this.min = append([]int{minElement}, this.min...)
}

func (this *MinStack) Pop() {
	this.vals = this.vals[1:]
	this.min = this.min[1:]
	if len(this.min) > 0 {
		minElement = this.min[0]
	} else {
		minElement = math.MaxInt
	}
}

func (this *MinStack) Top() int {
	return this.vals[0]
}

func (this *MinStack) GetMin() int {
	return this.min[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
