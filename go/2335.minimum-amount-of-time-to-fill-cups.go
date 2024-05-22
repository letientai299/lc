package main

// Minimum Amount of Time to Fill Cups
//
// # Easy
//
// https://leetcode.com/problems/minimum-amount-of-time-to-fill-cups
//
// You have a water dispenser that can dispense cold, warm, and hot water. Every
// second, you can either fill up `2` cups with **different** types of water, or
// `1` cup of any type of water.
//
// You are given a **0-indexed** integer array `amount` of length `3` where
// `amount[0]`, `amount[1]`, and `amount[2]` denote the number of cold, warm,
// and hot water cups you need to fill respectively. Return _the **minimum**
// number of seconds needed to fill up all the cups_.
//
// **Example 1:**
//
// ```
// Input: amount = [1,4,2]
// Output: 4
// Explanation: One way to fill up the cups is:
// Second 1: Fill up a cold cup and a warm cup.
// Second 2: Fill up a warm cup and a hot cup.
// Second 3: Fill up a warm cup and a hot cup.
// Second 4: Fill up a warm cup.
// It can be proven that 4 is the minimum number of seconds needed.
//
// ```
//
// **Example 2:**
//
// ```
// Input: amount = [5,4,4]
// Output: 7
// Explanation: One way to fill up the cups is:
// Second 1: Fill up a cold cup, and a hot cup.
// Second 2: Fill up a cold cup, and a warm cup.
// Second 3: Fill up a cold cup, and a warm cup.
// Second 4: Fill up a warm cup, and a hot cup.
// Second 5: Fill up a cold cup, and a hot cup.
// Second 6: Fill up a cold cup, and a warm cup.
// Second 7: Fill up a hot cup.
//
// ```
//
// **Example 3:**
//
// ```
// Input: amount = [5,0,0]
// Output: 5
// Explanation: Every second, we fill up a cold cup.
//
// ```
//
// **Constraints:**
//
// - `amount.length == 3`
// - `0 <= amount[i] <= 100`
func fillCups(amount []int) int {
	a, b, c := amount[0], amount[1], amount[2]
	res := 0
	for a > 0 && b > 0 && c > 0 {
		if a <= b && a <= c {
			a, c = c, a
		} else if b <= a && b <= c {
			b, c = c, b
		}

		a--
		b--
		res++
	}

	if a == 0 {
		res += max(b, c)
	} else if b == 0 {
		res += max(a, c)
	} else {
		res += max(a, b)
	}

	return res
}

// TODO (tai): cna we find an O(1) solution for this?
