// Leetcode problem 739 - Daily Temperatures
// https://leetcode.com/problems/daily-temperatures/

// Given an array of integers temperatures represents the daily temperatures,
// return an array answer such that answer[i] is the number of days you have
// to wait after the ith day to get a warmer temperature. If there is no
// future day for which this is possible, keep answer[i] == 0 instead.

// Input: temperatures = [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]

package leetcode739

// O(N^2) - Brute Force
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures)-1)
	for i, val := range temperatures[0 : len(temperatures)-1] {
		for j, comp := range temperatures[i+1:] {
			if val < comp {
				res[i] = j + 1
				break
			}
		}
	}
	return append(res, 0)
}

// O(N)
// https://leetcode.com/problems/daily-temperatures/discuss/1574808/C%2B%2BPython-3-Simple-Solutions-w-Explanation-Examples-and-Images-or-2-Monotonic-Stack-Approaches
