// 383. Ransom Note
// Easy

// Given two stings ransomNote and magazine, return true if ransomNote can be constructed from magazine and false otherwise.

// Each letter in magazine can only be used once in ransomNote.

// Example 1:

// Input: ransomNote = "a", magazine = "b"
// Output: false
// Example 2:

// Input: ransomNote = "aa", magazine = "ab"
// Output: false
// Example 3:

// Input: ransomNote = "aa", magazine = "aab"
// Output: true

// Constraints:

// 1 <= ransomNote.length, magazine.length <= 105
// ransomNote and magazine consist of lowercase English letters.

package main

import "fmt"

func main() {
	rs := "a"
	m := "b"
	fmt.Println(canConstruct(rs, m))
}

func canConstruct(ransomNote string, magazine string) bool {
	countArr := make([]int, 26)
	for _, rr := range magazine {
		countArr[rr-'a']++
	}
	for _, rm := range ransomNote {
		countArr[rm-'a']--
		if countArr[rm-'a'] < 0 {
			return false
		}
	}
	return true
}
