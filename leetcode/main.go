package main

import (
	"fmt"

	arrays "github.com/guerraglucas/Coding-Challenges/leetcode/arrays"
	mockAss "github.com/guerraglucas/Coding-Challenges/leetcode/mockAssessment"
)

func main() {
	fmt.Println(arrays.MergeAlternately("ilv ", " oeu"))
	fmt.Println(arrays.GcdOfStrings("ABCABCABC", "ABC"))
	fmt.Println(arrays.KidsWithCandies([]int{1, 2, 0, 4, 5}, 2))
	fmt.Println(mockAss.FindPeakElement([]int{-2147483648, -2147483647}))
	fmt.Println(mockAss.FindPeakElement([]int{1, 2}))
	fmt.Println(mockAss.RepeatedSubstringPattern("ababababababab"))
}
