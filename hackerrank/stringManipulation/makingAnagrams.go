// A student is taking a cryptography class and has found anagrams to be very useful. Two strings are
// anagrams of each other if the first string's letters can be rearranged to form the second string.
//
// In other words, both strings must contain the same exact letters in the same exact frequency. For example,
// bacdc and dcbac are anagrams, but bacdc and dcbad are not.

// The student decides on an encryption scheme that involves two large strings. The encryption is dependent on
//  the minimum number of character deletions required to make the two strings anagrams. Determine this number.

// Given two strings, a and b, that may or may not be of the same length, determine the minimum number of
//  character deletions required to make a and b anagrams. Any characters can be deleted from either of the strings.

// Example
// a = 'cde'
// b = 'dcf'
// Delete e from a and f from b so that the remaining strings are cd and dc which are anagrams. This takes 2 character deletions.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'makeAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func makeAnagram(a string, b string) int32 {
	mapStringA := make(map[rune]int)
	mapStringB := make(map[rune]int)
	sumOfDeletions := int32(0)

	for _, rune := range a {
		mapStringA[rune]++
	}
	for _, rune := range b {
		mapStringB[rune]++
	}

	for key, value := range mapStringA {
		if _, ok := mapStringB[key]; ok {
			if value == mapStringB[key] {
				continue
			}
			sumOfDeletions += int32(returnAbsoluteValue(value, mapStringB[key]))
		} else {
			sumOfDeletions += int32(value)
		}
	}
	for key, value := range mapStringB {
		if _, ok := mapStringA[key]; ok {
			if value == mapStringA[key] {
				continue
			}
			sumOfDeletions += int32(returnAbsoluteValue(value, mapStringA[key]))
		} else {
			sumOfDeletions += int32(value)
		}
	}
	return sumOfDeletions
}

func returnAbsoluteValue(a, b int) int {
	if a-b <= 0 {
		return 0
	}
	return a - b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

	fmt.Fprintf(writer, "%d\n", res)

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
