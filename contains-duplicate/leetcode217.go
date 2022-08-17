package leetcode217

// Time O(n^2)
// Space O(1)
func containsDuplicateBruteForce(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i += 1 {
		curr := i
		for j := i + 1; j < n; j += 1 {
			if nums[curr] == nums[j] {
				return true
			}
		}
	}
	return false
}

// Time O(n)
// Space O(n)
type set map[int]bool

func containsDuplicate(nums []int) bool {
	set := make(set, 0)

	// This is O(1) because for slices Go gets the length from the cache
	// n := len(nums)

	for i := 0; i < len(nums); i += 1 {
		if set[nums[i]] {
			return true
		}
		set[nums[i]] = true
	}

	return false
}
