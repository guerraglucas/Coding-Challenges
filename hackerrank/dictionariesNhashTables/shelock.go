// Two strings are anagrams of each other if the letters of one string can be rearranged to form the other string.
// Given a string, find the number of pairs of substrings of the string that are anagrams of each other.
// For example s = mom, the list of all anagrammatic pairs is [m,m], [mo,om] at positions [[0],[2]], [[0,1],[1,2]] respectively.
// Function Description
// Complete the function sherlockAndAnagrams in the editor below.
// sherlockAndAnagrams has the following parameter(s):
// string s: a string
// Returns
// int: the number of unordered anagrammatic pairs of substrings in s
// Input Format
// The first line contains an integer q, the number of queries.
// Each of the next q lines contains a string s to analyze.
// Constraints
// 1 <= q <= 10
// 2 <= |s| <= 100
// String s contains only lowercase letters ascii[a-z].
// Sample Input 0
// 2
// abba
// abcd
// Sample Output 0
// 4
// 0
// Explanation 0
// The list of all anagrammatic pairs is [a,a], [ab,ba], [b,b] and [abb,bba] at positions [[0],[3]], [[0,1],[2,3]], [[1],[2]] and [[0,1,2],[1,2,3]] respectively.
// No anagrammatic pairs exist in the second query as no character repeats.
// Sample Input 1
// 2
// ifailuhkqq
// kkkk
// Sample Output 1
// 3
// 10
// Explanation 1
// For the first query, we have anagram pairs [i,i], [q,q] and [ifa,fai] at positions [[0],[3]], [[8],[9]] and [[0,1,2],[1,2,3]] respectively.
// For the second query:
// There are 6 anagrams of the form [k,k] at positions [[0],[1]], [[0],[2]], [[0],[3]], [[1],[2]], [[1],[3]] and [[2],[3]].
// There are 3 anagrams of the form [kk,kk] at positions [[0,1],[1,2]], [[0,1],[2,3]] and [[1,2],[2,3]].

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'sherlockAndAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func sherlockAndAnagrams(s string) int32 {

	// initiate the variables responsible for storing the map of anagrams and the counter
	mapOfSubStrings := make(map[string]int)
	count := 0

	// iterate trough the string, checking all possible SORTED substrings
	for i := 0; i < len(s); i++ {
		// resets the substring each time the i is incremented
		subString := ""

		// iterate to check all possible substring from i index through the end of the remaining string
		for j := i; j < len(s); j++ {

			// concatenates the current j index into the substring
			subString += string(s[j])

			// sorts the substring, so it can be stored into the map and be easily comparable to the future substrings, as both will be sorted in ascending order
			sortedSubstring := sortSubstring(subString)

			// increments the count of the sorted substring key on the map, or creates a key:value if the map didn't already have the current sorted substring
			mapOfSubStrings[sortedSubstring]++
		}
	}

	// after generating the map of all possible substrings, handle the calculation of anagrams.
	// this is done by getting each count of subStrings in the map, multiplying it by (count - 1), this is important because
	// in the case that a substring doesn't have a anagram pair, the count will be 1
	// as such, it will read: count += 1 * 0 / 2, which results zero and dont add up to the overall count of anagram pairs
	for _, anagramCount := range mapOfSubStrings {
		count += anagramCount * (anagramCount - 1) / 2
	}

	return int32(count)
}

// sorts the substrings by comparing the rune values
func sortSubstring(subString string) string {

	// generate a slice of all the substring runes
	sliceOfChars := []rune(subString)

	// execute the slice method from the sort package to sort in ascending order, comparing the rune values
	sort.Slice(sliceOfChars, func(i, j int) bool {
		return sliceOfChars[i] < sliceOfChars[j]
	})

	return string(sliceOfChars)
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

		result := sherlockAndAnagrams(s)

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
