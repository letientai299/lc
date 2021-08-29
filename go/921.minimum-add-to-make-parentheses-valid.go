package main

// Minimum Add to Make Parentheses Valid
//
// Medium
//
// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
//
// A parentheses string is valid if and only if:
//
// - It is the empty string,
// - It can be written as `AB` (`A` concatenated with `B`), where `A` and `B`
// are valid strings, or
// - It can be written as `(A)`, where `A` is a valid string.
//
// You are given a parentheses string `s`. In one move, you can insert a
// parenthesis at any position of the string.
//
// - For example, if `s = "()))"`, you can insert an opening parenthesis to be
// `"(()))"` or a closing parenthesis to be `"())))"`.
//
// Return _the minimum number of moves required to make_ `s` _valid_.
//
// **Example 1:**
//
// ```
// Input: s = "())"
// Output: 1
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "((("
// Output: 3
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "()"
// Output: 0
//
// ```
//
// **Example 4:**
//
// ```
// Input: s = "()))(("
// Output: 4
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 1000`
// - `s[i]` is either `'('` or `')'`.
func minAddToMakeValid(s string) int {
	r := 0
	deep := 0
	for _, c := range s {
		if c == ')' {
			if deep > 0 {
				deep--
			} else {
				r++
			}
		} else {
			deep++
		}
	}
	return r + deep
}

// Good for interview
