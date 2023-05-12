// You are given an unordered array consisting of consecutive integers  [1, 2, 3, ..., n] without any duplicates.
// You are allowed to swap any two elements.
// You need to find the minimum number of swaps required to sort the array in ascending order.
// For example, given the array  we perform the following steps:
// Example:
// arr=[7, 1, 3, 2, 4, 5, 6]

// Perform the following steps:
// i   arr                         swap (indices)
// 0   [7, 1, 3, 2, 4, 5, 6]   swap (0,3)
// 1   [2, 1, 3, 7, 4, 5, 6]   swap (0,1)
// 2   [1, 2, 3, 7, 4, 5, 6]   swap (3,4)
// 3   [1, 2, 3, 4, 7, 5, 6]   swap (4,5)
// 4   [1, 2, 3, 4, 5, 7, 6]   swap (5,6)
// 5   [1, 2, 3, 4, 5, 6, 7]

// It took 5 swaps to sort the array.
// Function Description
// Complete the function minimumSwaps in the editor below.
// minimumSwaps has the following parameter(s):
// int arr[n]: an unordered array of integers
// Returns
// int: the minimum number of swaps to sort the array
// Input Format
// The first line contains an integer, n , the size of arr.
// The second line contains n space-separated integers .

// Sample Input 0
// 4
// 4 3 1 2
// Sample Output 0
// 3
// Explanation 0
// Given array arr: [4,3,1,2]
// After swapping (0,2) we get arr: [1,3,4,2]
// After swapping (1,2) we get arr: [1,4,3,2]
// After swapping (1,3) we get arr: [1,2,3,4]
// So, we need a minimum of 3 swaps to sort the array in ascending order.

// Sample Input 1
// 5
// 2 3 4 1 5
// Sample Output 1
// 3
// Explanation 1
// Given array arr: [2,3,4,1,5]
// After swapping (2,3) we get arr: [2,3,1,4,5]
// After swapping (0,1) we get arr: [3,2,1,4,5]
// After swapping (0,2) we get arr: [1,2,3,4,5]
// So, we need a minimum of 3 swaps to sort the array in ascending order.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func minimumSwaps(arr []int32) int32 {
	minSwaps := int32(0)
	// start from the left
	for i := int32(0); i < int32(len(arr)); i++ {
		// check if the position of the element matches the element value,
		// if so jumps to next iteration
		if arr[i] == i+1 {
			continue
		}

		// iterate for the next elements and checks which matches the i positon desired
		// value.
		// if so, calls the swap function that swap those elements and increment the counter
		for j := int32(i + 1); j < int32(len(arr)); j++ {
			if arr[j] == i+1 {
				swap(arr, i, j, &minSwaps)
				continue
			}
		}
	}
	return minSwaps

}

// takes the array, two indexes to swap and the pointer to the counter, to increment it
func swap(arr []int32, i, j int32, count *int32) {
	arr[i], arr[j] = arr[j], arr[i]
	*count++
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

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
