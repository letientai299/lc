package main

// Largest Odd Number in String
//
// Easy
//
// https://leetcode.com/problems/largest-odd-number-in-string/
//
// You are given a string `num`, representing a large integer. Return _the
// **largest-valued odd** integer (as a string) that is a **non-empty
// substring** of_ `num` _, or an empty string_`""` _if no odd integer exists_.
//
// A **substring** is a contiguous sequence of characters within a string.
//
// **Example 1:**
//
// ```
// Input: num = "52"
// Output: "5"
// Explanation: The only non-empty substrings are "5", "2", and "52". "5" is the
// only odd number.
//
// ```
//
// **Example 2:**
//
// ```
// Input: num = "4206"
// Output: ""
// Explanation: There are no odd numbers in "4206".
//
// ```
//
// **Example 3:**
//
// ```
// Input: num = "35427"
// Output: "35427"
// Explanation: "35427" is already an odd number.
//
// ```
//
// **Constraints:**
//
// - `1 <= num.length <= 105`
// - `num` only consists of digits and does not contain any leading zeros.
func largestOddNumber(num string) string {
	i := len(num) - 1
	for ; i >= 0; i-- {
		if (num[i]-'0')%2 == 1 {
			break
		}
	}

	if i < 0 {
		return ""
	}

	return num[:i+1]
}
