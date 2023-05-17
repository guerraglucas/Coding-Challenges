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
	fmt.Println(arrays.CanPlaceFlowers([]int{0, 0, 1, 0, 0}, 2))
	fmt.Println(mockAss.FindPeakElement([]int{-2147483648, -2147483647}))
	fmt.Println(mockAss.FindPeakElement([]int{1, 2}))
	fmt.Println(mockAss.RepeatedSubstringPattern("ababababababab"))
	fmt.Println(mockAss.Fib(6))
	fmt.Println(mockAss.DistributeCandies([]int{2, 4, 5, 12, 1, 3, 4, 5, 6, 8}))
}
