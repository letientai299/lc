package main

// Maximum Ascending Subarray Sum
//
// Easy
//
// https://leetcode.com/problems/maximum-ascending-subarray-sum/
//
// Given an array of positive integers `nums`, return the _maximum possible sum
// of an **ascending** subarray in_ `nums`.
//
// A subarray is defined as a contiguous sequence of numbers in an array.
//
// A subarray `[numsl, numsl+1, ..., numsr-1, numsr]` is **ascending** if for
// all `i` where `l <= i < r`, `numsi  < numsi+1`. Note that a subarray of size
// `1` is **ascending**.
//
// **Example 1:**
//
// ```
// Input: nums = [10,20,30,5,10,50]
// Output: 65
// Explanation: [5,10,50] is the ascending subarray with the maximum sum of 65.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [10,20,30,40,50]
// Output: 150
// Explanation: [10,20,30,40,50] is the ascending subarray with the maximum sum
// of 150.
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [12,17,15,13,10,11,12]
// Output: 33
// Explanation: [10,11,12] is the ascending subarray with the maximum sum of 33.
//
// ```
//
// **Example 4:**
//
// ```
// Input: nums = [100,10,1]
// Output: 100
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 100`
// - `1 <= nums[i] <= 100`
func maxAscendingSum(nums []int) int {
	max := 0
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if cur > max {
				max = cur
			}
			cur = nums[i]
			continue
		}

		cur += nums[i]
	}

	if cur > max {
		max = cur
	}

	return max
}
