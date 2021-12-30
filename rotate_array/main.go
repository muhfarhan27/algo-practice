// 189. Rotate Array
// Medium

// Given an array, rotate the array to the right by k steps, where k is non-negative.

// Example 1:

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
// Example 2:

// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

package main

import "fmt"

func main() {
	var (
		nums = []int{1, 2, 3, 4, 5, 6, 7}
		k    = 3
	)
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	k = k % len(nums)
	cap := len(nums) - 1 - k
	temp := []int{}
	last := len(nums) - 1
	var i = 0
	for i = last; i > cap; i-- {
		temp = append(temp, nums[i])
	}
	left := 0
	right := len(temp) - 1
	for left <= right {
		temp[left], temp[right] = temp[right], temp[left]
		left = left + 1
		right = right - 1
	}
	for index, num := range nums {
		if index > cap {
			break
		}
		temp = append(temp, num)
	}
	copy(nums, temp)
}
