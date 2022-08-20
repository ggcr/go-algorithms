package leetcode1

// BRUTE FORCE
// Time: O(N^2)
// Space: O(1)
func twoSum(nums []int, target int) []int {
	for i, val_i := range nums {
		want := target - val_i
		if want >= 0 {
			for j, val_j := range nums[i+1:] {
				if val_j == want {
					return []int{i, j + i + 1}
				}
			}
		}
	}
	return []int{}
}

// Using a hashmap
// Time: O(N)
// Space: O(N)
func twoSumMap(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))
	for i, v := range nums {
		want := target - v
		if _, ok := hashmap[want]; ok {
			return []int{hashmap[want], i}
		} else {
			hashmap[v] = i
		}
	}
	return nil
}
