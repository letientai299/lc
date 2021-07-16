package main

// Maximum Population Year
//
// Easy
//
// https://leetcode.com/problems/maximum-population-year/
//
// You are given a 2D integer array `logs` where each `logs[i] = [birthi,
// deathi]` indicates the birth and death years of the `ith` person.
//
// The **population** of some year `x` is the number of people alive during that
// year. The `ith` person is counted in year `x`'s population if `x` is in the
// **inclusive** range `[birthi, deathi - 1]`. Note that the person is **not**
// counted in the year that they die.
//
// Return _the **earliest** year with the **maximum population**_.
//
// **Example 1:**
//
// ```
// Input: logs = [[1993,1999],[2000,2010]]
// Output: 1993
// Explanation: The maximum population is 1, and 1993 is the earliest year with
// this population.
//
// ```
//
// **Example 2:**
//
// ```
// Input: logs = [[1950,1961],[1960,1971],[1970,1981]]
// Output: 1960
// Explanation:
// The maximum population is 2, and it had happened in years 1960 and 1970.
// The earlier year between them is 1960.
// ```
//
// **Constraints:**
//
// - `1 <= logs.length <= 100`
// - `1950 <= birthi < deathi <= 2050`
func maximumPopulation(logs [][]int) int {
	var years [101]int
	const start = 1950
	for _, l := range logs {
		years[l[0]-start]++
		years[l[1]-start]--
	}

	maxID := -1
	max := -1
	cur := 0
	for i := 0; i < len(years); i++ {
		cur += years[i]
		if cur > max {
			maxID = i
			max = cur
		}
	}
	return maxID + start
}
