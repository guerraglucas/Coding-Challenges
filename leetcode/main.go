package main

import (
	"fmt"

	// arrays "github.com/guerraglucas/Coding-Challenges/leetcode/arrays"
	mockAss "github.com/guerraglucas/Coding-Challenges/leetcode/mockAssessment"
	twoPointers "github.com/guerraglucas/Coding-Challenges/leetcode/twoPointers"
)

func main() {
	// fmt.Println(arrays.MergeAlternately("ilv ", " oeu"))
	// fmt.Println(arrays.GcdOfStrings("ABCABCABC", "ABC"))
	// fmt.Println(arrays.KidsWithCandies([]int{1, 2, 0, 4, 5}, 2))
	// fmt.Println(arrays.CanPlaceFlowers([]int{0, 0, 1, 0, 0}, 2))
	// fmt.Println(arrays.ReverseVowels("hello"))
	// fmt.Println(arrays.ReverseWords("   get ready   for the   next   battle   "))
	// fmt.Println(mockAss.FindPeakElement([]int{-2147483648, -2147483647}))
	// fmt.Println(mockAss.FindPeakElement([]int{1, 2}))
	// fmt.Println(mockAss.RepeatedSubstringPattern("ababababababab"))
	// fmt.Println(mockAss.Fib(6))
	// fmt.Println(mockAss.DistributeCandies([]int{2, 4, 5, 12, 1, 3, 4, 5, 6, 8}))
	// printTree(mockAss.TrimBST(&mockAss.TreeNode{Val: 3, Left: &mockAss.TreeNode{Val: 0, Left: nil, Right: &mockAss.TreeNode{Val: 2, Left: &mockAss.TreeNode{Val: 1}}}, Right: &mockAss.TreeNode{Val: 4}}, 1, 3))
	// printTree(mockAss.TrimBST(&mockAss.TreeNode{Val: 1, Left: &mockAss.TreeNode{Val: 0, Left: nil, Right: nil}, Right: &mockAss.TreeNode{Val: 2}}, 1, 2))
	// fmt.Println(mockAss.OrangesRotting([][]int{{2, 0, 1}, {1, 2, 1}, {0, 0, 1}}))
	// fmt.Println(mockAss.RemoveElement([]int{2, 4, 41, 1, 2, 3, 4}, 2))
	// fmt.Println(mockAss.CombinationSum([]int{2, 3, 6, 7}, 7))
	twoPointers.MoveZeroes([]int{0, 0, 1})

}

func printTree(node *mockAss.TreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	printTree(node.Left)
	printTree(node.Right)
}
