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

func GcdOfStrings(str1 string, str2 string) string {
	lenA := len(str1)
	lenB := len(str2)

	// now we need to find the greatest common divisor between the two lengths
	gcd := gdcCalculator(lenA, lenB)

	// get the substring by slice of the str1 until the gcd index
	substring := str1[:gcd]

	// if concatenating str1 and str2 is equal to the inverse (str2 + str1) AND the substring is equal to the slice of the str2 until
	// index of gcd, means that the substring is the greatest divider of both string
	if str1+str2 == str2+str1 && substring == str2[:gcd] {
		return substring
	} else {
		return ""
	}

}

// function gets the greates common divisor by recursively calling itself to check if the remaining of the division of elements is == 0
// which means that the 'a' parameter is the greatest common divisor
func gdcCalculator(a, b int) int {
	if b == 0 {
		return a
	}
	return gdcCalculator(b, a%b)
}
