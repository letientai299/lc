package main

import "strings"

// Determine if String Halves Are Alike
//
// Easy
//
// https://leetcode.com/problems/determine-if-string-halves-are-alike/
//
// You are given a string `s` of even length. Split this string into two halves
// of equal lengths, and let `a` be the first half and `b` be the second half.
//
// Two strings are **alike** if they have the same number of vowels (`'a'`,
// `'e'`, `'i'`, `'o'`, `'u'`, `'A'`, `'E'`, `'I'`, `'O'`, `'U'`). Notice that
// `s` contains uppercase and lowercase letters.
//
// Return `true` _if_ `a` _and_ `b` _are **alike**_. Otherwise, return `false`.
//
// **Example 1:**
//
// ```
// Input: s = "book"
// Output: true
// Explanation: a = "bo" and b = "ok". a has 1 vowel and b has 1 vowel.
// Therefore, they are alike.
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "textbook"
// Output: false
// Explanation: a = "text" and b = "book". a has 1 vowel whereas b has 2.
// Therefore, they are not alike.
// Notice that the vowel o is counted twice.
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "MerryChristmas"
// Output: false
//
// ```
//
// **Example 4:**
//
// ```
// Input: s = "AbCdEfGh"
// Output: true
//
// ```
//
// **Constraints:**
//
// - `2 <= s.length <= 1000`
// - `s.length` is even.
// - `s` consists of **uppercase and lowercase** letters.
func halvesAreAlike(s string) bool {
	n := len(s)
	diff := 0
	for i := 0; i < n/2; i++ {
		a, b := s[i], s[n-i-1]
		if strings.IndexByte("aeiouAEIOU", a) >= 0 {
			diff++
		}
		if strings.IndexByte("aeiouAEIOU", b) >= 0 {
			diff--
		}
	}
	return diff == 0
}
