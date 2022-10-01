package leetcode217

/* ####################################################

	Brute force
	- Time O(n^2)
	- Space O(n)

#################################################### */

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

/* ####################################################

	Merge sort method
	- Time O(n logn)
	- Space O(n)

#################################################### */

func getMedian(nums []int) (index int) {
	n := len(nums)
	if len(nums)%2 == 1 {
		n -= 1
	}
	return n / 2
}

func mergeSort(left, right []int) ([]int, bool) {
	cond := false

	arr := make([]int, len(left)+len(right))

	// Merge the two subarrays
	i := 0
	j := 0
	k := 0

	for i < len(left) && j < len(right) {
		if left[i] == right[j] {
			cond = true
		}
		if left[i] <= right[j] {
			arr[k] = left[i]
			i += 1
		} else {
			arr[k] = right[j]
			j += 1
		}
		k += 1
	}

	for i < len(left) {
		arr[k] = left[i]
		i += 1
		k += 1
	}

	for j < len(right) {
		arr[k] = right[j]
		j += 1
		k += 1
	}

	return arr, cond
}

func containsDuplicateMergeSort(nums []int) ([]int, bool) {
	if len(nums) < 2 {
		return nums, false
	}

	// Split into left and right
	m := getMedian(nums)
	left := nums[:m]
	right := nums[m:]

	L, cond := containsDuplicateMergeSort(left)
	R, cond := containsDuplicateMergeSort(right)

	res, cond := mergeSort(L, R)

	return res, cond
}

/* ####################################################

	Hash-set method
	- Time O(n)
	- Space O(n)

#################################################### */

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
