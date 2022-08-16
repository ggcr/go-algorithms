// Leetcode problem 61- Rotate List
// https://leetcode.com/problems/wiggle-sort-ii/

// Given an integer array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....

package leetcode324

import (
	"sort"
)

// utils

func checkWiggle(nums []int) bool {
	prev := nums[0]
	var current int
	var next int
	for id, val := range nums[1 : len(nums)-1] {
		current = val
		next = nums[id+2]
		if prev <= current && current <= next {
			return false
		} else if prev >= current && current >= next {
			return false
		}
		prev = current
	}
	return true
}

func getMedian(nums []int) int {
	len := len(nums)
	if len%2 != 0 {
		len = len - 1
	}
	return len / 2
}

func wiggleSort(nums []int) {
	arr := make([]int, len(nums))
	copy(arr, nums)
	// Sort array
	sorted_arr := sort.IntSlice(arr)
	// Reverse it
	sort.Sort(sort.Reverse(sorted_arr))
	// Find Upper Median position
	median := getMedian(nums)
	// Higher-End
	H := sorted_arr[0:median]
	// Lower-End
	L := sorted_arr[median:]

	for i, _ := range nums {
		if i%2 == 0 {
			nums[i] = L[i/2]
		} else {
			nums[i] = H[(i-1)/2]
		}
	}
}
