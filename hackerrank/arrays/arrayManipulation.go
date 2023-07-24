// Starting with a 1-indexed array of zeros and a list of operations, for each operation add a value to each the array element between two given indices, inclusive. Once all operations have been performed, return the maximum value in the array.
// Example
// n = 10
// queries = [[1,5,3],[4,8,7],[6,9,1]]

// Queries are interpreted as follows:
//     a b k
//     1 5 3
//     4 8 7
//     6 9 1
// Add the values of k between the indices a and b inclusive:
// index->	 1 2 3  4  5 6 7 8 9 10
// 	[0,0,0, 0, 0,0,0,0,0, 0]
// 	[3,3,3, 3, 3,0,0,0,0, 0]
// 	[3,3,3,10,10,7,7,7,0, 0]
// 	[3,3,3,10,10,8,8,8,1, 0]
// The largest value is 10 after all operations are performed.
// Function Description
// Complete the function arrayManipulation in the editor below.
// arrayManipulation has the following parameters:
// int n - the number of elements in the array
// int queries[q][3] - a two dimensional array of queries where each queries[i] contains three integers, a, b, and k.
// Returns
// int - the maximum value in the resultant array
// Input Format
// The first line contains two space-separated integers n and m, the size of the array and the number of operations.
// Each of the next m lines contains three space-separated integers a,b and k, the left index, right index and summand.
// Constraints
// 3 <= n <= 10^7
// 1 <= m <= 2*10^5
// 1 <= a <= b <= n
// 0 <= k <= 10^9

// Sample Input
// 5 3
// 1 2 100
// 2 5 100
// 3 4 100
// Sample Output
// 200
// Explanation
// After the first update the list is 100 100 0 0 0.
// After the second update list is 100 200 100 100 100.
// After the third update list is 100 200 200 200 100.
// The maximum value is 200.

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
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func arrayManipulation(n int32, queries [][]int32) int64 {
	// Write your code here

	// create the base array of zeros
	baseSlice := make([]int64, n)
	var maxValue, prefixSum int64

	// iterate through the queries array and make the operations
	for _, query := range queries {
		valueToAdd := int64(query[2])
		startingIndex := query[0] - 1
		endingIndex := query[1]

		baseSlice[startingIndex] += valueToAdd
		if endingIndex < n {
			baseSlice[endingIndex] -= valueToAdd
		}
	}
	fmt.Println(baseSlice)
	for _, value := range baseSlice {
		prefixSum += value
		if prefixSum > maxValue {
			maxValue = prefixSum
		}
	}

	return maxValue

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

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

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
