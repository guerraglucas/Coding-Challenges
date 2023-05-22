// 151. Reverse Words in a String
// Medium
// 5.9K
// 4.6K
// Companies
// Given an input string s, reverse the order of the words.

// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

// Return a string of the words in reverse order concatenated by a single space.

// Note that s may contain leading or trailing spaces or multiple spaces between two words.
//The returned string should only have a single space separating the words. Do not include any extra spaces.

// Example 1:

// Input: s = "the sky is blue"
// Output: "blue is sky the"
// Example 2:

// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.
// Example 3:

// Input: s = "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

package leetcodeArrayChallenges

func ReverseWords(s string) string {
	var arrSubstring []string
	var subString string

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			if subString != "" {

				arrSubstring = append(arrSubstring, subString)
				subString = ""
			}
			continue
		}

		subString += string(s[i])
	}
	if subString != "" {
		arrSubstring = append(arrSubstring, subString)
	}

	return reverseArrString(arrSubstring)

}

func reverseArrString(arr []string) string {
	var reversedSubstring string
	for i := len(arr) - 1; i >= 0; i-- {

		if arr[i] != "" {
			if i == 0 {

				reversedSubstring += arr[i]
			} else {
				reversedSubstring += arr[i] + " "
			}
		}
	}
	return reversedSubstring
}
