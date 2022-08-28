// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

package leetcode128

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hashmap := make(map[int]int, len(nums))
	res := 0
	val := 0
	max := 0
	for _, value := range nums {
		hashmap[value] = 1

	}
	for value, _ := range hashmap {
		if hashmap[value] != 1 {
			hashmap[value] = 1
			if _, ok_plus := hashmap[value+1]; ok_plus {
				val = value + 1
				for {
					if _, ok := hashmap[val]; ok {
						hashmap[val] = 1
						res += 1
						val += 1
						continue
					}
					break
				}
			}
			if _, ok_prev := hashmap[value-1]; ok_prev {
				val = value - 1
				for {
					if _, ok := hashmap[val]; ok {
						res += 1
						val -= 1
						continue
					}
					break
				}
			}
			if res > max {
				max = res
			}
		}
		res = 0
	}
	return max + 1
}
