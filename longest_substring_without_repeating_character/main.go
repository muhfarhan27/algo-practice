// 3. Longest Substring Without Repeating Characters
// Medium

// Given a string s, find the length of the longest substring without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Constraints:

// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.

package main

import "fmt"

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	charCounter := make(map[rune]int)
	l := 0
	globalMax := 0
	for r, runeS := range s {
		charCounter[runeS]++
		for charCounter[runeS] > 1 && l < r {
			charCounter[rune(s[l])]--
			l++
		}
		globalMax = max(globalMax, r-l+1)
	}
	return globalMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
