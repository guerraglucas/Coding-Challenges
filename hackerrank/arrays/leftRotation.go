// A left rotation operation on an array shifts each of the array's elements 1 unit to the left.
// For example, if 2 left rotations are performed on array [1,2,3,4,5], then the array would become [3,4,5,1,2].
// Given an array a of n integers and a number, d, perform d left rotations on the array.
// Return the updated array to be printed as a single line of space-separated integers.
// Function Description
// Complete the function rotLeft in the editor below. It should return the resulting array of integers.
// rotLeft has the following parameter(s):
// An array of integers a.
// An integer d, the number of rotations.
// Input Format
// The first line contains two space-separated integers n and d, the size of a and the number of left rotations you must perform.
// The second line contains n space-separated integers a[i].
// Constraints
// 1 <= n <= 10^5
// 1 <= d <= n
// 1 <= a[i] <= 10^6
// Output Format
// Print a single line of n space-separated integers denoting the final state of the array after performing d left rotations.
// Sample Input
// 5 4
// 1 2 3 4 5
// Sample Output
// 5 1 2 3 4
// Explanation
// When we perform d = 4 left rotations, the array undergoes the following sequence of changes:
// [1,2,3,4,5] -> [2,3,4,5,1] -> [3,4,5,1,2] -> [4,5,1,2,3] -> [5,1,2,3,4]

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
 * Complete the 'rotLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER d
 */

func rotLeft(a []int32, d int32) []int32 {
	// Write your code here

	// modulus operand to reduce the times the for loop iterates, if d > len(a)
	d = d % int32(len(a))

	for i := int32(0); i < d; i++ {
		a = append(a[1:], a[0])
	}
	return a
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := rotLeft(a, d)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
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
