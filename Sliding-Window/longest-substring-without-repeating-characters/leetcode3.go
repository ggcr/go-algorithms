package leetcode3

// This was tricky because the leading pointer is actually our reference
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max := 0
	out := 0

	ptr_in := 0
	ptr_fi := 1

	cond := false
	for ptr_in < len(s) {
		cond = false
		for i := ptr_in; i < ptr_fi; i++ {
			if s[i] == s[ptr_fi] {
				out = ptr_fi - ptr_in
				if out > max {
					max = out
				}
				ptr_in = i + 1
				cond = true
				break
			}
		}
		if !cond {
			out = ptr_fi - ptr_in + 1
			if out > max {
				max = out
			}
			if ptr_fi == len(s)-1 {
				ptr_in += 1
			} else {
				ptr_fi += 1
			}
		}
		if ptr_in == ptr_fi {
			if ptr_fi == len(s)-1 {
				return max
			}
			ptr_fi += 1
		}
	}
	return max
}
