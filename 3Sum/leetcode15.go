// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]

package leetcode15

func checkMap(val_i, val_j, val_k int, m map[int][]int) (res bool) {
	for _, vals_i := range m[val_i] {
		for _, vals_j := range m[val_j] {
			for _, vals_k := range m[val_k] {
				if vals_j == vals_k && vals_i == vals_j {
					return true
				}
			}
		}
	}
	return false
}

func appendMap(val_i, val_j, val_k int, m map[int][]int, global_ctr int) map[int][]int {
	m[val_i] = append(m[val_i], global_ctr)
	m[val_j] = append(m[val_j], global_ctr)
	m[val_k] = append(m[val_k], global_ctr)
	return m
}

// O(N^3)
// Brute force, three pointers and checking each pointer
func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 1 && nums[0] == 0 {
		return append(res, []int{0, 0, 0})
	}

	m := make(map[int][]int)
	global_ctr := 0
	allZeros := false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if !allZeros && (nums[i] == 0 && nums[j] == 0 && nums[k] == 0) {
					m = appendMap(nums[i], nums[j], nums[k], m, global_ctr)
					global_ctr += 1
					res = append(res, []int{nums[i], nums[j], nums[k]})
					allZeros = true
				} else if nums[i]+nums[j]+nums[k] == 0 && !checkMap(nums[i], nums[j], nums[k], m) {
					m = appendMap(nums[i], nums[j], nums[k], m, global_ctr)
					global_ctr += 1
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return res
}
