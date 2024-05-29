package main

// Alternating Digit Sum
//
// # Easy
//
// https://leetcode.com/problems/alternating-digit-sum
//
// You are given a positive integer `n`. Each digit of `n` has a sign according
// to the following rules:
//
// - The **most significant digit** is assigned a **positive** sign.
// - Each other digit has an opposite sign to its adjacent digits.
//
// Return _the sum of all digits with their corresponding sign_.
//
// **Example 1:**
//
// ```
// Input: n = 521
// Output: 4
// Explanation: (+5) + (-2) + (+1) = 4.
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 111
// Output: 1
// Explanation: (+1) + (-1) + (+1) = 1.
//
// ```
//
// **Example 3:**
//
// ```
// Input: n = 886996
// Output: 0
// Explanation: (+8) + (-8) + (+6) + (-9) + (+9) + (-6) = 0.
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 109`
func alternateDigitSum(n int) int {
	s := 0
	sig := 1
	cnt := 0
	for n > 0 {
		d := n % 10
		n /= 10
		s += sig * d
		sig *= -1
		cnt++
	}
	if cnt%2 == 0 {
		s *= -1
	}

	return s
}
