package main

// Count Common Words With One Occurrence
//
// Easy
//
// https://leetcode.com/problems/count-common-words-with-one-occurrence/
//
// Given two string arrays `words1` and `words2`, return _the number of strings
// that appear **exactly once** in **each** of the two arrays._
//
// **Example 1:**
//
// ```
// Input: words1 = ["leetcode","is","amazing","as","is"], words2 =
// ["amazing","leetcode","is"]
// Output: 2
// Explanation:
// - "leetcode" appears exactly once in each of the two arrays. We count this
// string.
// - "amazing" appears exactly once in each of the two arrays. We count this
// string.
// - "is" appears in each of the two arrays, but there are 2 occurrences of it
// in words1. We do not count this string.
// - "as" appears once in words1, but does not appear in words2. We do not count
// this string.
// Thus, there are 2 strings that appear exactly once in each of the two arrays.
//
// ```
//
// **Example 2:**
//
// ```
// Input: words1 = ["b","bb","bbb"], words2 = ["a","aa","aaa"]
// Output: 0
// Explanation: There are no strings that appear in each of the two arrays.
//
// ```
//
// **Example 3:**
//
// ```
// Input: words1 = ["a","ab"], words2 = ["a","a","a","ab"]
// Output: 1
// Explanation: The only string that appears exactly once in each of the two
// arrays is "ab".
//
// ```
//
// **Constraints:**
//
// - `1 <= words1.length, words2.length <= 1000`
// - `1 <= words1[i].length, words2[j].length <= 30`
// - `words1[i]` and `words2[j]` consists only of lowercase English letters.
func countWords(words1 []string, words2 []string) int {
	m1 := make(map[string]int)
	for _, w := range words1 {
		m1[w]++
	}

	m2 := make(map[string]int)
	for _, w := range words2 {
		m2[w]++
	}

	r := 0
	for w, v := range m1 {
		if v == 1 && m2[w] == 1 {
			r++
		}
	}
	return r
}
