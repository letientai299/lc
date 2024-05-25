package main

// Remove Letter To Equalize Frequency
//
// # Easy
//
// https://leetcode.com/problems/remove-letter-to-equalize-frequency
//
// You are given a **0-indexed** string `word`, consisting of lowercase English
// letters. You need to select **one** index and **remove** the letter at that
// index from `word` so that the **frequency** of every letter present in `word`
// is equal.
//
// Return`true` _if it is possible to remove one letter so that the frequency of
// all letters in_ `word` _are equal, and_ `false` _otherwise_.
//
// **Note:**
//
// - The **frequency** of a letter `x` is the number of times it occurs in the
// string.
// - You **must** remove exactly one letter and cannot choose to do nothing.
//
// **Example 1:**
//
// ```
// Input: word = "abcc"
// Output: true
// Explanation: Select index 3 and delete it: word becomes "abc" and each
// character has a frequency of 1.
//
// ```
//
// **Example 2:**
//
// ```
// Input: word = "aazz"
// Output: false
// Explanation: We must delete a character, so either the frequency of "a" is 1
// and the frequency of "z" is 2, or vice versa. It is impossible to make all
// present letters have equal frequency.
//
// ```
//
// **Constraints:**
//
// - `2 <= word.length <= 100`
// - `word` consists of lowercase English letters only.
func equalFrequency(word string) bool {
	m := make(map[rune]int)
	for _, c := range word {
		m[c]++
	}

	a, b := -1, -1
	cntA, cntB := 0, 0
	for _, v := range m {
		if a == -1 {
			a = v
			cntA = 1
			continue
		}

		if v != a && b == -1 {
			b = v
			cntB = 1
			continue
		}

		if v == a {
			cntA++
			continue
		}

		if v == b {
			cntB++
			continue
		}

		return false
	}

	return (a == b-1 && cntB == 1) ||
		(b == a-1 && cntA == 1) ||
		(b == -1 && a == 1) ||
		(b == -1 && cntA == 1) ||
		(a == 1 && cntA == 1) ||
		(b == 1 && cntB == 1)
}
