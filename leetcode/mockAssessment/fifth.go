// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

// Consider the number of elements in nums which are not equal to val be k, to get accepted,
// you need to do the following things:

// Change the array nums such that the first k elements of nums contain the elements which are not
// equal to val. The remaining elements of nums are not important as well as the size of nums.
// Return k.
// Custom Judge:

// The judge will test your solution with the following code:

// int[] nums = [...]; // Input array
// int val = ...; // Value to remove
// int[] expectedNums = [...]; // The expected answer with correct length.
//                             // It is sorted with no values equaling val.

// int k = removeElement(nums, val); // Calls your implementation

// assert k == expectedNums.length;
// sort(nums, 0, k); // Sort the first k elements of nums
// for (int i = 0; i < actualLength; i++) {
//     assert nums[i] == expectedNums[i];
// }
// If all assertions pass, then your solution will be accepted.

// Example 1:

// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 2.
// It does not matter what you leave beyond the returned k (hence they are underscores).
// Example 2:

// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
// Note that the five elements can be returned in any order.
// It does not matter what you leave beyond the returned k (hence they are underscores).

// Constraints:

// 0 <= nums.length <= 100
// 0 <= nums[i] <= 50
// 0 <= val <= 100

package mockAssessment

import (
	"sort"
)

func RemoveElement(nums []int, val int) int {
	counter := len(nums)
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = -1
			counter--
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return counter
}

// Given an array of distinct integers candidates and a target integer target, return a list of all unique
// combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

// The same number may be chosen from candidates an unlimited number of times. Two combinations are unique
// if the frequency of at least one of the chosen numbers is different.

// The test cases are generated such that the number of unique combinations that sum up to target is less than 150
// combinations for the given input.

// Example 1:

// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.
// Example 2:

// Input: candidates = [2,3,5], target = 8
// Output: [[2,2,2,2],[2,3,3],[3,5]]
// Example 3:

// Input: candidates = [2], target = 1
// Output: []

// Constraints:

// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// All elements of candidates are distinct.
// 1 <= target <= 40

func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var combination []int

	// Sort candidates to handle duplicates and ensure ascending order in combinations
	sort.Ints(candidates)

	backtrack(&result, &combination, candidates, target, 0, 0)

	return result
}

func backtrack(result *[][]int, combination *[]int, candidates []int, target, currentSum, startIdx int) {
	if currentSum == target {
		// Add a copy of the current combination to the result
		tmp := make([]int, len(*combination))
		copy(tmp, *combination)
		*result = append(*result, tmp)
		return
	}

	if currentSum > target {
		return
	}

	for i := startIdx; i < len(candidates); i++ {
		*combination = append(*combination, candidates[i])                              // Choose the current candidate
		backtrack(result, combination, candidates, target, currentSum+candidates[i], i) // Recursively explore with the current candidate
		*combination = (*combination)[:len(*combination)-1]                             // Remove the current candidate from the combination
	}
}
