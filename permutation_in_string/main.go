// 567. Permutation in String
// Medium

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

// In other words, return true if one of s1's permutations is the substring of s2.

// Example 1:

// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").

// Example 2:

// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

// Constraints:

// 1 <= s1.length, s2.length <= 104
// s1 and s2 consist of lowercase English letters.

package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	j := len(s1)
	s1Arr := getLetterArray(s1)
	for i := 0; i < len(s2); i++ {
		if j > len(s2) {
			return false
		}
		subStr := s2[i:j]
		if compareArraySameLen(s1Arr, getLetterArray(subStr)) {
			return true
		}
		j++
	}
	return false
}

func getLetterArray(s string) []int {
	res := make([]int, 26)
	for _, char := range s {
		res[char-'a']++
	}
	return res
}

func compareArraySameLen(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
