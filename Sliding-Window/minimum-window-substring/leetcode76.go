package leetcode76

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	copy_strMap := make(map[string]int)
	for _, v := range t {
		copy_strMap[string(v)] += 1
	}

	ptr_in := 0
	ptr_fi := 0
	count := 0
	res := ""
	max := ""
	shrink := false

	for ptr_in <= len(s)-len(t) {
		if count == len(t) {
			if len(res) < len(max) || len(max) == 0 {
				max = res
			}
		}

		if shrink {
			// Reduce res string
			if len(t) == count {
				ptr_in += 1
				if len(res) < 2 {
					return max
				}
				val := res[0]
				res = res[1:]
				if v, ok := copy_strMap[string(val)]; ok {
					copy_strMap[string(val)] += 1
					if v == 0 {
						count -= 1
					}
				}
			} else {
				if _, ok := copy_strMap[string(s[ptr_in])]; ok {
					if ptr_fi != len(s)-1 {
						shrink = false
						ptr_fi += 1
					} else {
						return max
					}
				} else {
					ptr_in += 1
					res = res[1:]
				}
			}
		} else {
			// Grow res string
			res += string(s[ptr_fi])
			if v, ok := copy_strMap[string(s[ptr_fi])]; ok {
				copy_strMap[string(s[ptr_fi])] -= 1
				if v > 0 {
					count += 1
				}
			}
			if ptr_fi == len(s)-1 || count == len(t) {
				shrink = true
			} else {
				ptr_fi += 1
			}
		}

	}
	return max
}
