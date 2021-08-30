package main

import "sort"

// Minimum Increment to Make Array Unique
//
// Medium
//
// https://leetcode.com/problems/minimum-increment-to-make-array-unique/
//
// You are given an integer array `nums`. In one move, you can pick an index `i`
// where `0 <= i < nums.length` and increment `nums[i]` by `1`.
//
// Return _the minimum number of moves to make every value in_ `nums`
// _**unique**_.
//
// **Example 1:**
//
// ```
// Input: nums = [1,2,2]
// Output: 1
// Explanation: After 1 move, the array could be [1, 2, 3].
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [3,2,1,2,1,7]
// Output: 6
// Explanation: After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
// It can be shown with 5 or less moves that it is impossible for the array to
// have all unique values.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 105`
// - `0 <= nums[i] <= 105`
func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	r := 0
	for i := 1; i < len(nums); i++ {
		a, b := nums[i-1], nums[i]
		if b > a {
			continue
		}
		r += a - b + 1
		nums[i] = a + 1
	}
	return r
}

// TODO (tai): can be faster, sort result in 264 ms, 23.08%.
