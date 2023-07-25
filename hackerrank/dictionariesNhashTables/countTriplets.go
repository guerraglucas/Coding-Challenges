// You are given an array and you need to find number of tripets of indices (i, j, k) such that the elements at those indices
// are in geometric progression for a given common ratio r and i < j < k.
// For example, arr = [1, 4, 16, 64]. If r = 4, we have [1, 4, 16] and [4, 16, 64] at indices (0, 1, 2) and (1, 2, 3).
// Function Description
// Complete the countTriplets function in the editor below. It should return the number of triplets forming a
// geometric progression for a given r as an integer.
// countTriplets has the following parameter(s):
// arr: an array of integers
// r: an integer, the common ratio
// Input Format
// The first line contains two space-separated integers n and r, the size of arr and the common ratio.
// The next line contains n space-seperated integers arr[i].
// Constraints
// 1 <= n <= 10^5
// 1 <= r <= 10^9
// 1 <= arr[i] <= 10^9
// Output Format
// Return the count of triplets that form a geometric progression.
// Sample Input 0
// 4 2
// 1 2 2 4
// Sample Output 0
// 2
// Explanation 0
// There are 2 triplets in satisfying our criteria, whose indices are (0, 1, 3) and (0, 2, 3)
// Sample Input 1
// 6 3
// 1 3 9 9 27 81
// Sample Output 1
// 6
// Explanation 1
// The triplets satisfying are index (0, 1, 2), (0, 1, 3), (1, 2, 4), (1, 3, 4), (2, 4, 5) and (3, 4, 5).
// Sample Input 2
// 5 5
// 1 5 5 25 125
// Sample Output 2
// 4
// Explanation 2
// The triplets satisfying are index (0, 1, 3), (0, 2, 3), (1, 3, 4), (2, 3, 4).

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	// start variable to store the count
	var count int64 = 0
	// create two maps to store the left and right values
	// those maps will store the value as the key and the count as the value
	var left, right map[int64]int64
	left = make(map[int64]int64)
	right = make(map[int64]int64)
	// loop through the array and store the values in the right map
	for _, v := range arr {
		right[v]++
	}
	// loop through the array again
	// decrement the right value
	// if the value is divisible by r
	// add the left value * right value to the count
	// increment the left value
	for _, v := range arr {
		right[v]--
		if v%r == 0 {
			count += left[v/r] * right[v*r]
		}
		left[v]++
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
