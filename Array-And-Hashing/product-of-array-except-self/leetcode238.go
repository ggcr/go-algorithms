// Given an integer array nums, return an array answer such that answer[i]
// is equal to the product of all the elements of nums except nums[i].

// You must write an algorithm that runs in O(n) time and without using the
// division operation.

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

package leetcode238

// Time: O(N)
// Space: O(N)
func productExceptSelf(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(nums); i += 1 {
		arr[i] = arr[i-1] * nums[i]
	}

	reverse_arr := make([]int, len(nums))
	reverse_arr[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i -= 1 {
		reverse_arr[i] = reverse_arr[i+1] * nums[i]
	}

	nums[0] = reverse_arr[1]
	nums[len(nums)-1] = arr[len(nums)-2]
	for i := 1; i < len(nums)-1; i += 1 {
		nums[i] = arr[len(arr)-1] - ((arr[i] - arr[i-1]) * reverse_arr[i+1])
	}

	return nums
}
