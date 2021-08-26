package main

import (
	"sort"
)

// Friends Of Appropriate Ages
//
// Medium
//
// https://leetcode.com/problems/friends-of-appropriate-ages/
//
// There are `n` persons on a social media website. You are given an integer
// array `ages` where `ages[i]` is the age of the `ith` person.
//
// A Person `x` will not send a friend request to a person `y` (`x != y`) if any
// of the following conditions is true:
//
// - `age[y] <= 0.5 * age[x] + 7`
// - `age[y] > age[x]`
// - `age[y] > 100 && age[x] < 100`
//
// Otherwise, `x` will send a friend request to `y`.
//
// Note that if `x` sends a request to `y`, `y` will not necessarily send a
// request to `x`. Also, a person will not send a friend request to themself.
//
// Return _the total number of friend requests made_.
//
// **Example 1:**
//
// ```
// Input: ages = [16,16]
// Output: 2
// Explanation: 2 people friend request each other.
//
// ```
//
// **Example 2:**
//
// ```
// Input: ages = [16,17,18]
// Output: 2
// Explanation: Friend requests are made 17 -> 16, 18 -> 17.
//
// ```
//
// **Example 3:**
//
// ```
// Input: ages = [20,30,100,110,120]
// Output: 3
// Explanation: Friend requests are made 110 -> 100, 120 -> 110, 120 -> 100.
//
// ```
//
// **Constraints:**
//
// - `n == ages.length`
// - `1 <= n <= 2 * 104`
// - `1 <= ages[i] <= 120`
func numFriendRequests(ages []int) int {
	i := 0
	m := make(map[int]int)
	for _, v := range ages {
		m[v]++
		if m[v] == 1 {
			ages[i] = v
			i++
		}
	}

	ages = ages[:i]
	sort.Ints(ages)

	counts := make([]int, i)
	pre := 0
	for i, v := range ages {
		pre += m[v]
		counts[i] = pre
	}

	// TODO (tai): 2-pointers is faster

	res := 0
	for i, x := range ages {
		y := x/2 + 7

		v := m[x]
		if x > y {
			res += v * (v - 1)
		}

		if i == 0 {
			continue
		}

		l := sort.SearchInts(ages, y+1)
		if l >= i {
			continue
		}

		if l == 0 {
			res += counts[i-1] * v
		} else {
			res += (counts[i-1] - counts[l-1]) * v
		}
	}

	return res
}
