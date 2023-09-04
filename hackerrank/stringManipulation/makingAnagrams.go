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
	// create a map for each string
	mapA := make(map[rune]int32)
	mapB := make(map[rune]int32)
	// create a map for the result
	resultMap := make(map[rune]int32)
	for _, char := range a {
		mapA[char]++
	}
	for _, char := range b {
		mapB[char]++
	}
	// iterate over the first map
	for key, countA := range mapA {
		// if the key is in the second map, check the value
		countB, exists := mapB[key]
		if exists {
			// Character exists in both strings
			resultMap[key] = abs(countA - countB)
		} else {
			// if the key is not in the second map, add the value to the result map
			// Character only exists in string a
			resultMap[key] = countA
		}
	}
	// iterate over the second map
	for key, countB := range mapB {
		_, exists := mapA[key]
		// if the key is in the first map, check the value
		if !exists {
			// if the key is not in the first map, add the value to the result map
			// Character only exists in string b
			resultMap[key] = countB
		}
	}
	// iterate over the result map and sum the values
	// return the sum
	var sum int32
	for _, value := range resultMap {
		sum += value
	}
	return sum
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
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
