// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

package leetcode125

import (
	"strings"
)

// O(n)
func strip(s string) string {
	s = strings.ToLower(s)
	var str_b strings.Builder
	for i := 0; i < len(s); i += 1 {
		b := s[i]
		if 'a' <= b && b <= 'z' ||
			'0' <= b && b <= '9' {
			str_b.WriteByte(b)
		}
	}
	return str_b.String()
}

// O(n)
func isPalindrome(s string) bool {
	s = strip(s)

	var median int = len(s) / 2

	i := median
	k := median
	if len(s)%2 == 0 {
		k -= 1
	}

	for i < len(s) && k >= 0 {
		if s[i] != s[k] {
			return false
		}
		i += 1
		k -= 1
	}

	return true
}
