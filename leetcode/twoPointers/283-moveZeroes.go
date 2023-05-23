// 283. Move Zeroes
// Easy
// 13.6K
// 344
// Companies
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// Example 1:

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:

// Input: nums = [0]
// Output: [0]

// Constraints:

// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1

// Follow up: Could you minimize the total number of operations done?

package twoPointersLeetCode

import "fmt"

func MoveZeroes(nums []int) {
	left := 0
	right := left + 1

	for right < len(nums) {
		if nums[left] == 0 && nums[right] == 0 {
			right++
			continue
		}
		if nums[left] == 0 && nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			fmt.Println(nums)
		}
		left++
		right++
	}
}
