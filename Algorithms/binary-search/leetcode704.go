package leetcode704

func binarySearch(nums []int, i, target int) int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
		return -1
	}

	half := len(nums) / 2
	modulo := half % 2

	if nums[half] == target {
		return i
	}

	if nums[half] > target {
		return binarySearch(nums[:half], i-(half/2)-(modulo), target)
	} else {
		return binarySearch(nums[half+modulo:], i+(half/2)+(modulo), target)
	}

}

func search(nums []int, target int) int {
	res := binarySearch(nums, len(nums)/2, target)
	return res
}
