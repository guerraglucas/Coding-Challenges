// There is a string, , of lowercase English letters that is repeated infinitely many times.
//  Given an integer, , find and print the number of letter a's in the first  letters of the infinite string.
// Example
// s = 'abcac'
// n = 10
// The substring we consider is abcacabcac, the first 10 characters of the infinite string. There are 4 occurrences of a in the substring.

// Function Description
// Complete the repeatedString function in the editor below.
// repeatedString has the following parameter(s):
// s: a string to repeat
// n: the number of characters to consider
// Returns
// int: the frequency of a in the substring
// Input Format

// The first line contains a single string, s.
// The second line contains an integer, n.

package solution

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	// Write your code here
	mapOfChars := make(map[string]int)

	for _, char := range s {
		mapOfChars[string(char)]++
	}

	// count checks how many times the 's' string FULLY fits inside the repeated string with 'n' length
	// and gets the amount of 'a' that are present on this repeated string
	count := int64(mapOfChars["a"]) * (n / int64(len(s)))

	// now we iterate only over the remaining letters to be added, which is subpart of the original 's' string
	remainingLetters := n % int64(len(s))
	for i := int64(0); i < remainingLetters; i++ {
		if s[i] == 'a' {
			count++
		}
	}

	return count
}
