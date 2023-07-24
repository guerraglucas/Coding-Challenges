// Given two strings, determine if they share a common substring. A substring may be as small as one character.
// Example
// s1 = 'and'
// s2 = 'art'
// These share the common substring a.
// s1 = 'be'
// s2 = 'cat'
// These do not share a substring.
// Function Description
// Complete the function twoStrings in the editor below.
// twoStrings has the following parameter(s):
// string s1: a string
// string s2: another string
// Returns
// string: either YES or NO
// Input Format
// The first line contains a single integer p, the number of test cases.
// The following p pairs of lines are as follows:
// The first line contains string s1.
// The second line contains string s2.
// Constraints
// s1 and s2 consist of characters in the range ascii[a-z].
// 1 <= p <= 10
// 1 <= |s1|, |s2| <= 10^5
// Output Format
// For each pair of strings, return YES or NO.
// Sample Input
// 2
// hello
// world
// hi
// world
// Sample Output
// YES
// NO
// Explanation
// We have p = 2 pairs to check:
// s1 = 'hello', s2 = 'world'. The substrings 'o' and 'l' are common to both strings.
// a = 'hi', b = 'world'. s1 and s2 share no common substrings.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'twoStrings' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func twoStrings(s1 string, s2 string) string {
	// Write your code here
	mapS1 := make(map[string]int)

	for _, char := range s1 {
		mapS1[string(char)]++
	}
	for _, char := range s2 {
		if _, ok := mapS1[string(char)]; ok {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s1 := readLine(reader)

		s2 := readLine(reader)

		result := twoStrings(s1, s2)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
