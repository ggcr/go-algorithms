package leetcode15

func twoSum(numbers []int, target int) []int {
	if len(numbers) <= 2 {
		return []int{1, 2}
	}
	ptr_in := 0
	ptr_fi := len(numbers) - 1
	for ptr_in < len(numbers) && ptr_fi >= 0 {
		if numbers[ptr_in]+numbers[ptr_fi] > target {
			ptr_fi -= 1
		} else if numbers[ptr_in]+numbers[ptr_fi] < target {
			ptr_in += 1
		}
		if numbers[ptr_in]+numbers[ptr_fi] == target {
			return []int{ptr_in + 1, ptr_fi + 1}
		}
	}
	return []int{}
}
