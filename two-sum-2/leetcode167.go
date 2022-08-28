// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]

// If we compare Two Sum I and Two Sum II, this time we a re given the input array SORTED, meaning that we do not need a hashmap anymore, we just have to loop the array in O(N) and return the sum.

package leetcode167

func twoSum(numbers []int, target int) []int {
	// i will be the behind pointer
	// k will be the advanced pointer
	if len(numbers) <= 2 {
		return []int{1, 2}
	}
	i := 0
	k := 1
	for i < len(numbers) {
		if k >= len(numbers) ||
			numbers[i]+numbers[k] > target {
			i += 1
			k = i + 1
			continue
		} else if numbers[i]+numbers[k] == target {
			return []int{i + 1, k + 1}
		}
		k += 1
	}
	return []int{}
}
