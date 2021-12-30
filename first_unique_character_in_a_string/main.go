// 387. First Unique Character in a String
// Easy

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

// Example 1:

// Input: s = "leetcode"
// Output: 0
// Example 2:

// Input: s = "loveleetcode"
// Output: 2
// Example 3:

// Input: s = "aabb"
// Output: -1

// Constraints:

// 1 <= s.length <= 105
// s consists of only lowercase English letters.

package main

import (
	"fmt"
	"math"
)

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for i, rs := range s {
		if _, exist := m[rs]; !exist {
			m[rs] = i
		} else {
			m[rs] = -1
		}
	}
	min := math.MaxInt32
	for _, val := range m {
		if val >= 0 {
			min = minimum(min, val)
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	return min
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
