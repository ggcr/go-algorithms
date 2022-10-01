package leetcode121

// Sliding window O(N)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	ptr_in := 0
	ptr_fi := 1

	out := 0
	max := 0

	for ptr_in < ptr_fi && ptr_in < len(prices) {
		if prices[ptr_in] < prices[ptr_fi] {
			out = prices[ptr_fi] - prices[ptr_in]
			if out > max {
				max = out
			}
			if ptr_fi == len(prices)-1 {
				ptr_in += 1
			} else {
				ptr_fi += 1
			}
		} else {
			if ptr_in == ptr_fi-1 && ptr_fi != len(prices)-1 {
				ptr_fi += 1
			}
			ptr_in += 1
		}
	}
	return max
}
