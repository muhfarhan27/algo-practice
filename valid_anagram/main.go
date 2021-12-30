// 242. Valid Anagram
// Easy

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:

// Input: s = "rat", t = "car"
// Output: false

// Constraints:

// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.

// Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	countArr := make([]int, 26)
	for _, runeS := range s {
		countArr[runeS-'a']++
	}
	for _, runeT := range t {
		countArr[runeT-'a']--
	}
	for _, value := range countArr {
		if value != 0 {
			return false
		}
	}
	return true
}
