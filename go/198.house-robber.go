package main

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 *
 * https://leetcode.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (40.84%)
 * Total Accepted:    298.8K
 * Total Submissions: 731.7K
 * Testcase Example:  '[1,2,3,1]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed, the only constraint stopping
 * you from robbing each of them is that adjacent houses have security system
 * connected and it will automatically contact the police if two adjacent
 * houses were broken into on the same night.
 *
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 * Example 2:
 *
 *
 * Input: [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house
 * 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 *
 *
 */
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	robs := make([]int, n)
	dp[0] = nums[0]
	robs[0] = 0

	dp[1] = nums[1]
	robs[1] = 1
	if nums[0] > nums[1] {
		dp[1] = nums[0]
		robs[1] = 0
	}

	for i := 2; i < n; i++ {

		get := nums[i] + dp[i-2]

		if robs[i-1] != i-1 && dp[i-1] > dp[i-2] {
			get = nums[i] + dp[i-1]
			dp[i] = get
			robs[i] = i
			continue
		}

		// maintain current profit
		if get < dp[i-1] {
			dp[i] = dp[i-1]
			robs[i] = robs[i-1]
			continue
		}

		dp[i] = get
		robs[i] = i
	}

	return dp[n-1]
}
