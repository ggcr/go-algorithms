package leetcode424

func characterReplacement(s string, k int) int {
	i := 0
	counts := make([]int, 26)
	maxCount := 0
	res := 0
	for j := 0; j < len(s); j++ {
		counts[s[j]-'A']++
		maxCount = max(maxCount, counts[s[j]-'A'])

		for i <= j && j-i+1-maxCount > k {
			counts[s[i]-'A']--
			i++
		}

		res = max(res, j-i+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
