// Given two strings s and t, return true if t is an anagram of s, and false
// otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a
// different word or phrase, typically using all the original letters
// exactly once.

// Input: s = "anagram", t = "nagaram"
// Output: true

package leetcode242

import (
	"reflect"
	"sort"
	"strings"
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

// This method is valid for all substrings t contained in s
// Time: O(n) => O(s+t)
// Space: O(n) => O(s+t)
func isSubAnagram(s string, t string) bool {
	if len(s) < len(t) {
		return false
	}
	hashset := make(map[byte]int)
	for id, val := range []byte(t) {
		t_val := hashset[val] + 1
		s_val := hashset[s[id]] - 1

		if val != s[id] {

			hashset[s[id]] = s_val
			hashset[val] = t_val
			if t_val == 0 {
				delete(hashset, val)
			}
			if s_val == 0 {
				delete(hashset, s[id])
			}
		}

	}
	return len(hashset) == 0
}

// Sort method
// Time: O(nlogn)
// Space: O(1) ???? O(n)
func isAnagramSort(s string, t string) bool {
	s_sort := strings.Split(s, "")
	t_sort := strings.Split(t, "")
	sort.Strings(s_sort)
	sort.Strings(t_sort)
	return reflect.DeepEqual(s_sort, t_sort)
}
