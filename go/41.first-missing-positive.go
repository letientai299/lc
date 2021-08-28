package main

// First Missing Positive
//
// Hard
//
// https://leetcode.com/problems/first-missing-positive/
//
// Given an unsorted integer array `nums`, return the smallest missing positive
// integer.
//
// You must implement an algorithm that runs in `O(n)` time and uses constant
// extra space.
//
// **Example 1:**
//
// ```
// Input: nums = [1,2,0]
// Output: 3
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [3,4,-1,1]
// Output: 2
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [7,8,9,11,12]
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 5 * 105`
// - `-231 <= nums[i] <= 231 - 1`
func firstMissingPositive(nums []int) int {
	n := len(nums)
	potential := n + 1
	seen := make(map[int]bool)
	for _, v := range nums {
		if v <= 0 {
			potential--
			continue
		}

		if v >= potential {
			potential--
			continue
		}

		if !seen[v] {
			seen[v] = true
		} else {
			potential--
		}
	}
	for i := 1; i <= potential; i++ {
		if !seen[i] {
			return i
		}
	}
	return potential
}

// TODO (tai): this is slow as hell, 342 ms, 5%
