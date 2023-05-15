// For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).

// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

// Example 1:

// Input: str1 = "ABCABC", str2 = "ABC"
// Output: "ABC"
// Example 2:

// Input: str1 = "ABABAB", str2 = "ABAB"
// Output: "AB"
// Example 3:

// Input: str1 = "LEET", str2 = "CODE"
// Output: ""
package leetcodeArrayChallenges

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func GcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	g := gcd(len1, len2)
	substr := str1[:g]
	if str1+str2 == str2+str1 && substr == str2[:g] {
		return substr
	}
	fmt.Println("Hello, playground")
	return ""
}
