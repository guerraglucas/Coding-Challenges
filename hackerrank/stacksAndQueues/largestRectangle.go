// Skyline Real Estate Developers is planning to demolish a number of old, unoccupied buildings and construct a shopping mall in their place.
// Your task is to find the largest solid area in which the mall can be constructed.
//
// There are a number of buildings in a certain two-dimensional landscape.
// Each building has a height, given by h[i] where i belongs to [1, n].
// If you join k adjacent buildings, they will form a solid rectangle of area k * min(h[i], h[i+1], ..., h[i+k-1]).
//
// For example, the heights array h = [3, 2, 3].
// A rectangle of height h = 2 and length k = 3 can be constructed within the boundaries.
// The area formed is h * k = 2 * 3 = 6.
//
// Function Description
// Complete the function largestRectangle int the editor below.
// It should return an integer representing the largest rectangle that can be formed within the bounds of consecutive buildings.
// largestRectangle has the following parameter(s):
// h: an array of integers representing building heights
//
// Returns
// int: the area of the largest rectangle that can be formed within the bounds of consecutive buildings
//
// Input Format
// The first line contains n, the number of buildings.
// The second line contains n space-separated integers, each representing the height of a building.
//
// Constraints
// 1 <= n <= 10^5
// 1 <= h[i] <= 10^6
//
// Output Format
// Print a long integer representing the maximum area of rectangle formed.
//
// Sample Input
// 5
// 1 2 3 4 5
//
// Sample Output
// 9
//
// Explanation
// An illustration of the test case follows.
// image
//
// The largest rectangle that can be formed within the boundaries of the consecutive buildings is 3 * 3 = 9.
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
 * Complete the 'largestRectangle' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY h as parameter.
 */

func largestRectangle(h []int32) int64 {
	// Write your code here
	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	hTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h []int32

	for i := 0; i < int(n); i++ {
		hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
		checkError(err)
		hItem := int32(hItemTemp)
		h = append(h, hItem)
	}

	result := largestRectangle(h)

	fmt.Fprintf(writer, "%d\n", result)

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
