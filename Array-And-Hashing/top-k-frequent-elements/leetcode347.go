// Given an integer array nums and an integer k, return
// the k most frequent elements. You may return the answer
// in any order.

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

package leetcode347

import (
	"sort"
)

// Time: O(N log N)
// Space: O(N)
func topKFrequentSort(nums []int, k int) []int {
	hashmap := make(map[int]int, len(nums))
	res := make([]int, k)
	for _, v := range nums {
		hashmap[v] += 1
	}

	type kv struct {
		Key   int
		Value int
	}

	var ss []kv
	for k, v := range hashmap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(x, y int) bool {
		return ss[x].Value > ss[y].Value
	})

	for i := 0; i < k; i += 1 {
		res[i] = ss[i].Key
	}

	return res
}

// Bucket Sort
// Time: O(N)
// Space: O(N)
func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int, len(nums))
	bucket := make([][]int, len(nums)+1)
	res := []int{}
	k_aux := k

	for _, v := range nums {
		hashmap[v] += 1
	}

	for key, v := range hashmap {
		bucket[v] = append(bucket[v], []int{key}...)
	}

	for ind := len(nums); ind >= 0; ind -= 1 {
		if len(bucket[ind]) >= k_aux {
			return append(res, bucket[ind][:k_aux]...)
		}
		res = append(res, bucket[ind]...)
		k_aux -= len(bucket[ind])
	}

	return nil
}
