// A bracket is considered to be any one of the following characters: (, ), {, }, [, or ].

// Two brackets are considered to be a matched pair if the an opening bracket (i.e., (, [, or {)
// occurs to the left of a closing bracket (i.e., ), ], or }) of the exact same type.
// There are three types of matched pairs of brackets: [], {}, and ().
// A matching pair of brackets is not balanced if the set of brackets it encloses are not matched.
// For example, {[(])} is not balanced because the contents in between { and } are not balanced.
// The pair of square brackets encloses a single, unbalanced opening bracket, (, and the pair of
// parentheses encloses a single, unbalanced closing square bracket, ].

// By this logic, we say a sequence of brackets is balanced if the following conditions are met:
// - It contains no unmatched brackets.
// - The subset of brackets enclosed within the confines of a matched pair of brackets is also a matched pair of brackets.

// Given n strings of brackets, determine whether each sequence of brackets is balanced.
// If a string is balanced, return YES. Otherwise, return NO.

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
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isBalanced(s string) string {
	// Write your code here
	var stack []string
	mapOfBrackets := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	for i := 0; i < len(s); i++ {
		// if the current element is an opening bracket, add to stack
		if _, ok := mapOfBrackets[string(s[i])]; ok {
			stack = append(stack, string(s[i]))
			continue
		}
		// the current element is a enclosing bracket
		// if the stack is currently empty or the current element is different from the expected enclosing bracket, return NO
		if len(stack) == 0 || string(s[i]) != mapOfBrackets[stack[len(stack)-1]] {
			return "NO"
		}
		// if the code reaches here, it means that the current element matches the expected enclosing bracket
		// remove the top element from the stack
		stack = stack[:len(stack)-1]
	}

	// confirms that at the end of the code the stack is empty, which is expected, then returns YES
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

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
