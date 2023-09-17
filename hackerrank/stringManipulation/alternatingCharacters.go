// You are given a string containing characters A and B only. Your task is to change it into a string such that there are no matching adjacent characters.
// To do this, you are allowed to delete zero or more characters in the string.
// Your task is to find the minimum number of required deletions.
// For example, given the string s = AABAAB, remove an A at positions 0 and 3 to make s = ABAB in 2 deletions.
//
// Function Description
// Complete the alternatingCharacters function in the editor below.
// alternatingCharacters has the following parameter(s):
// string s: a string
//
// Returns
// int: the minimum number of deletions required
//
// Input Format
// The first line contains an integer q, the number of queries.
// The next q lines each contain a string s.
//
// Constraints
// 1 <= q <= 10
// 1 <= |s| <= 10^5
// Each string s will consist only of characters A and B
//
// Sample Input
// 5
// AAAA
// BBBBB
// ABABABAB
// BABABA
// AAABBB
//
// Sample Output
// 3
// 4
// 0
// 0
// 4
//
// Explanation
// The characters marked red are the ones that can be deleted so that the string doesn't have matching consecutive characters.
// AAAA -> A (3 deletions)
// BBBBB -> B (4 deletions)
// ABABABAB -> ABABABAB (0 deletions)
// BABABA -> BABABA (0 deletions)
// AAABBB -> AB (4 deletions because to convert it to AB we must delete 2 A's and 2 B's)
//

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
 * Complete the 'alternatingCharacters' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func alternatingCharacters(s string) int32 {
	deletionCounter := int32(0)
	var charToCompare rune
	for i, char := range s {
		if i >= len(s) {
			break
		}
		if char == charToCompare {
			deletionCounter++
			i++
			continue
		}
		charToCompare = char
	}
	return deletionCounter
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
		s := readLine(reader)

		result := alternatingCharacters(s)

		fmt.Fprintf(writer, "%d\n", result)
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
