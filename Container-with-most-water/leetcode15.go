package leetcode11

func maxArea(height []int) int {
	max := 0
	area := 0

	ptr_in := 0
	ptr_fi := len(height) - 1

	global_ctr := len(height) - 1
	aux_mult := 0

	for ptr_in < ptr_fi {
		if height[ptr_in] >= height[ptr_fi] {
			aux_mult = height[ptr_fi]
			ptr_fi -= 1
		} else {
			aux_mult = height[ptr_in]
			ptr_in += 1
		}
		area = aux_mult * global_ctr
		if area > max {
			max = area
		}
		global_ctr -= 1
	}
	return max
}
