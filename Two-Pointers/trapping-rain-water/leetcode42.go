package leetcode42

// O(NlogN) Approach based on finding local max "valleys", minimum of 3 numbers where the extremas are bigger than the inner numbers

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findLocalMax(height []int, curr int) (int, int, int) {
	max := 0
	key := len(height) - 1
	total := 0
	ret := 0
	for k, v := range height[curr+1:] {
		if v > max {
			max = v
			key = k + curr + 1
			ret = total
			if max > height[curr] {
				return max, key, ret
			}
		}
		total += v
	}
	return max, key, ret
}

func trap(height []int) int {
	ptr_in := 0
	ptr_fi := 1
	sum := 0
	for ptr_in < ptr_fi && ptr_fi < len(height) {
		if height[ptr_in] > height[ptr_fi] {
			// Find other closest max
			_, key, total := findLocalMax(height, ptr_in)
			sum += (min(height[ptr_in], height[key]) * (key - ptr_in - 1)) - total
			ptr_fi = key + 1
			ptr_in = key
		} else {
			ptr_fi += 1
			ptr_in = ptr_fi - 1
		}
	}
	return sum
}
