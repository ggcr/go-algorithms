// Given an array of strings strs, group the anagrams together. You can return
// the answer in any order.

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

package leetcode49

import (
	"reflect"
)

// Comparing two hashmaps
// Time: O(n) => O(s+t)
// Space: O(n) => O(s+t)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashset_s := make(map[byte]int)
	hashset_t := make(map[byte]int)
	for id, val := range []byte(t) {
		hashset_s[s[id]] += 1
		hashset_t[val] += 1
	}
	return reflect.DeepEqual(hashset_s, hashset_t)
}

// BRUTE FORCE
// O(N^3)
func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}
	res := [][]string{}
	for _, val_i := range strs {
		check := false
		for j, val_j := range res {
			if len(val_j[0]) == len(val_i) {
				if isAnagram(val_j[0], val_i) {
					check = true
					res[j] = append(val_j, val_i)
				}
			}
		}
		if !check {
			res = append(res, []string{val_i})
		}
	}
	return res
}

// HASHMAP COUNT KEY STRATEGY
// Time: O(m * n)
// where m is the length of strs and n is avg length of each string inside

func groupAnagramsHashmap(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}
	res := [][]string{}
	hashmap := map[[26]int][]string{}
	for _, val_i := range strs {
		// count the nums
		ascii := [26]int{}
		for _, ascii_val := range val_i {
			ascii[ascii_val-97] += 1
		}
		// store it in the hashmap
		hashmap[ascii] = append(hashmap[ascii], val_i)
	}
	for _, val := range hashmap {
		res = append(res, val)
	}
	return res
}
