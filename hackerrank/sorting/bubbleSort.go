// Consider the following version of Bubble Sort:
//
// for (int i = 0; i < n; i++) {
//
//     for (int j = 0; j < n - 1; j++) {
//         // Swap adjacent elements if they are in decreasing order
//         if (a[j] > a[j + 1]) {
//             swap(a[j], a[j + 1]);
//         }
//     }
//
// }
// Given an array of integers, sort the array in ascending order using the Bubble Sort algorithm above. Once sorted, print the following three lines:
//
// 1 - Array is sorted in numSwaps swaps., where numSwaps is the number of swaps that took place.
// 2 - First Element: firstElement, where firstElement is the first element in the sorted array.
// 3 - Last Element: lastElement, where lastElement is the last element in the sorted array.
// Hint: To complete this challenge, you must add a variable that keeps a running tally of all swaps that occur during execution.
//
// For example, given a worst-case but small array to sort: a = [6, 4, 1] we go through the following steps:
//
// swap    a
// 0       [6, 4, 1]
// 1       [4, 6, 1]
// 2       [4, 1, 6]
// 3       [1, 4, 6]
// It took 3 swaps to sort the array. Output would be
//
// Array is sorted in 3 swaps.
// First Element: 1
// Last Element: 6
// Function Description
//
// Complete the function countSwaps in the editor below.
//
// countSwaps has the following parameter(s):
//
// int a[n]: an array of integers to sort
// Prints
//
// Print the three lines required, then return. No return value is expected.
//
// Input Format
//
// The first line contains an integer, n, the size of the array a.
// The second line contains n space-separated integers a[i].
//
// Constraints
//
// 2 <= n <= 600
// 1 <= a[i] <= 2 * 10^6

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
 * Complete the 'countSwaps' function below.
 *
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func countSwaps(a []int32) {
	// Write your code here
	swapCount := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapCount++
			}
		}
	}
	fmt.Printf("Array is sorted in %d swaps\n", swapCount)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d\n", a[len(a)-1])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	countSwaps(a)
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
