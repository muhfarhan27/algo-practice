// 350. Intersection of Two Arrays II
// Easy

// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

// Example 1:

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]
// Example 2:

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
// Explanation: [9,4] is also accepted.

// Constraints:

// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000

// Follow up:

// What if the given array is already sorted? How would you optimize your algorithm?
// What if nums1's size is small compared to nums2's size? Which algorithm is better?
// What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
	fmt.Println(intersectSort(nums1, nums2))
}

func intersectSort(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	i := 0
	j := 0
	temp := []int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			temp = append(temp, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return temp
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) > len(nums1) {
		return intersect(nums2, nums1)
	}
	numMap := make(map[int]int)
	for _, num := range nums1 {
		if _, exist := numMap[num]; !exist {
			numMap[num] = 1
		} else {
			numMap[num]++
		}
	}
	temp := []int{}
	for _, num := range nums2 {
		if _, exist := numMap[num]; exist {
			if numMap[num] > 0 {
				temp = append(temp, num)
				numMap[num]--
			}
		}
	}
	return temp
}
