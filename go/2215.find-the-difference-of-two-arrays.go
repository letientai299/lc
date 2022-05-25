package main

// Find the Difference of Two Arrays
//
// Easy
//
// https://leetcode.com/problems/find-the-difference-of-two-arrays/
//
// Given two **0-indexed** integer arrays `nums1` and `nums2`, return _a list_
// `answer` _of size_ `2` _where:_
//
// - `answer[0]` _is a list of all **distinct** integers in_ `nums1` _which are
// **not** present in_ `nums2` _._
// - `answer[1]` _is a list of all **distinct** integers in_ `nums2` _which are
// **not** present in_ `nums1`.
//
// **Note** that the integers in the lists may be returned in **any** order.
//
// **Example 1:**
//
// ```
// Input: nums1 = [1,2,3], nums2 = [2,4,6]
// Output: [[1,3],[4,6]]
// Explanation:
// For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1
// and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
// For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4
// and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].
// ```
//
// **Example 2:**
//
// ```
// Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
// Output: [[3],[]]
// Explanation:
// For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] ==
// nums1[3], their value is only included once and answer[0] = [3].
// Every integer in nums2 is present in nums1. Therefore, answer[1] = [].
//
// ```
//
// **Constraints:**
//
// - `1 <= nums1.length, nums2.length <= 1000`
// - `-1000 <= nums1[i], nums2[i] <= 1000`
func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]struct{})
	for _, n := range nums1 {
		m1[n] = struct{}{}
	}

	m2 := make(map[int]struct{})
	for _, n := range nums2 {
		m2[n] = struct{}{}
	}

	i := 0
	for n := range m1 {
		if _, ok := m2[n]; !ok {
			nums1[i] = n
			i++
		}
	}
	nums1 = nums1[:i]

	i = 0
	for n := range m2 {
		if _, ok := m1[n]; !ok {
			nums2[i] = n
			i++
		}
	}
	nums2 = nums2[:i]

	return [][]int{
		nums1, nums2,
	}
}
