// Given a string s, reverse only all the vowels in the string and return it.

// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

// Example 1:

// Input: s = "hello"
// Output: "holle"
// Example 2:

// Input: s = "leetcode"
// Output: "leotcede"

// Constraints:

// 1 <= s.length <= 3 * 105
// s consist of printable ASCII characters.
package leetcodeArrayChallenges

func ReverseVowels(s string) string {
	vowelIndices := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if isVowel(string(s[i])) {
			vowelIndices = append(vowelIndices, i)
		}
	}

	chars := []rune(s)
	// Reverse the vowels
	start := 0
	end := len(vowelIndices) - 1
	for start < end {
		chars[vowelIndices[start]], chars[vowelIndices[end]] = chars[vowelIndices[end]], chars[vowelIndices[start]]
		start++
		end--
	}

	return string(chars)
}

func isVowel(s string) bool {
	switch s {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		return true
	default:
		return false
	}
}
