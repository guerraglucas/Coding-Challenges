// Implement a function that takes in a string and encrypts it
// using the following algorithm:
//
// 1. Trim all spaces at the start and end of the string
// 2. Remove all the digits from 0 to 9
// 3. Reverse the string

// Example:
// "  a1bcd efg!h  " -> "hgfedcba"

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'ModifyString' function below and add imports if needed.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING str as parameter.
 */

func ModifyString(str string) string {
	reversedString := ""

	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			str = str[i+1:]
		} else {
			break
		}
	}
	fmt.Println(str)

	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == " " {
			str = str[:i]
		} else {
			break
		}
	}
	fmt.Println(str)

	for i := 0; i < len(str); i++ {
		if isNumerical(string(str[i])) {

			str = str[:i] + str[i+1:]
			i--

		}
	}
	fmt.Println(str)
	for i := len(str) - 1; i >= 0; i-- {
		reversedString += string(str[i])
	}
	return reversedString
}

func isNumerical(str string) bool {
	mapOfNumbers := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	if _, ok := mapOfNumbers[str]; ok {
		return true
	} else {
		return false
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	str := readLine(reader)

	result := ModifyString(str)

	fmt.Fprintf(writer, "%s\n", result)

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
