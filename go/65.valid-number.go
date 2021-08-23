package main

// Valid Number
//
// Hard
//
// https://leetcode.com/problems/valid-number/
//
// A **valid number** can be split up into these components (in order):
//
// 1. A **decimal number** or an **integer**.
// 2. (Optional) An `'e'` or `'E'`, followed by an **integer**.
//
// A **decimal number** can be split up into these components (in order):
//
// 1. (Optional) A sign character (either `'+'` or `'-'`).
// 2. One of the following formats:
//      1. One or more digits, followed by a dot `'.'`.
//     2. One or more digits, followed by a dot `'.'`, followed by one or more
// digits.
//     3. A dot `'.'`, followed by one or more digits.
//
// An **integer** can be split up into these components (in order):
//
// 1. (Optional) A sign character (either `'+'` or `'-'`).
// 2. One or more digits.
//
// For example, all the following are valid numbers: `["2", "0089", "-0.1",
// "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93",
// "-123.456e789"]`, while the following are not valid numbers: `["abc", "1a",
// "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]`.
//
// Given a string `s`, return `true` _if_ `s` _is a **valid number**_.
//
// **Example 1:**
//
// ```
// Input: s = "0"
// Output: true
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "e"
// Output: false
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "."
// Output: false
//
// ```
//
// **Example 4:**
//
// ```
// Input: s = ".1"
// Output: true
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 20`
// - `s` consists of only English letters (both uppercase and lowercase), digits
// (`0-9`), plus `'+'`, minus `'-'`, or dot `'.'`.
func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	dot := false
	digitBeforeExp := false
	exp := false
	digitAfterExp := false

	for i, c := range s {
		if '0' <= c && c <= '9' {
			if !exp {
				digitBeforeExp = true
			} else {
				digitAfterExp = true
			}
			continue
		}

		if c == '.' {
			if dot || exp {
				return false
			}
			dot = true
			continue
		}

		if c == '+' || c == '-' {
			if i > 0 && (s[i-1] == 'e' || s[i-1] == 'E') {
				continue
			}
			return false
		}

		if c == 'e' || c == 'E' {
			if exp {
				return false
			}
			exp = true
			continue
		}

		return false
	}

	return digitBeforeExp && (!exp || digitAfterExp)
}
