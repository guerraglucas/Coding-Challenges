// sort set of strings based on the following rules:
// 1. odd lenght string should precede an even lenght string
// 2 if both have odd lenghts, the shorter of the two should precede
// 3. if both have even lenghts, the longer of the two should precede
// 4. if the lenghts are the same, the strings should be in alphabetical order

//
// Example:
// Input: ["abc", "ab", "abcde", "a", "abcd"]
// Output: ["a", "abc", "abcde", "abcd", "ab"]

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
 * Complete the 'customSorting' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY strArr as parameter.
 */

func customSorting(strArr []string) []string {
	isOddLen := func(s string) bool {
		return len(s)%2 != 0
	}

	compare := func(a, b string) bool {
		if isOddLen(a) != isOddLen(b) {
			return isOddLen(a)
		} else if len(a) != len(b) {
			return len(a) < len(b)
		}
		return a < b
	}

	for i := 0; i < len(strArr); i++ {
		for j := i + 1; j < len(strArr); j++ {
			if compare(strArr[j], strArr[i]) {
				strArr[i], strArr[j] = strArr[j], strArr[i]
			}
		}
	}

	return strArr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	strArrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var strArr []string

	for i := 0; i < int(strArrCount); i++ {
		strArrItem := readLine(reader)
		strArr = append(strArr, strArrItem)
	}

	result := customSorting(strArr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
