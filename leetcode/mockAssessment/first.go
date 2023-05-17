// A peak element is an element that is strictly greater than its neighbors.

// Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

// You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.

// You must write an algorithm that runs in O(log n) time.

// Example 1:

// Input: nums = [1,2,3,1]
// Output: 2
// Explanation: 3 is a peak element and your function should return the index number 2.
// Example 2:

// Input: nums = [1,2,1,3,5,6,4]
// Output: 5
// Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.

package mockAssessment

import "strings"

func FindPeakElement(nums []int) int {
	indexToReturn := 0

	for i := 0; i < len(nums); i++ {
		var prev int
		var next int
		if i > 0 {
			prev = nums[i-1]
		}
		if i < len(nums)-1 {
			next = nums[i+1]
		}
		if i == 0 {
			prev = nums[i] - 1
		}
		if i == len(nums)-1 {
			next = nums[i] - 1
		}
		if nums[i] > prev && nums[i] > next {
			return i
		}
	}
	return indexToReturn
}

// Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.

// Example 1:

// Input: s = "abab"
// Output: true
// Explanation: It is the substring "ab" twice.
// Example 2:

// Input: s = "aba"
// Output: false
// Example 3:

// Input: s = "abcabcabcabc"
// Output: true
// Explanation: It is the substring "abc" four times or the substring "abcabc" twice.

func RepeatedSubstringPattern(s string) bool {
	substring := string(s[0])

	for i := 1; i < len(s); i++ {
		d := len(s) / len(substring)
		if s == strings.Repeat(substring, d) {
			return true
		}
		substring = substring + string(s[i])
	}

	return false
}
