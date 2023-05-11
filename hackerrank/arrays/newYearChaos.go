// It is New Year's Day and people are in line for the Wonderland rollercoaster ride. Each person
// wears a sticker indicating their initial position in the queue from 1 to n. Any person can bribe
// the person directly in front of them to swap positions, but they still wear their original sticker.
// One person can bribe at most two others. Determine the minimum number of bribes that took place to
// get to a given queue order. Print the number of bribes, or, if anyone has bribed more than two
// people, print Too chaotic.
// Example:
// q = [1,2,3,5,4,6,7,8]
// If person 5 bribes person 4, the queue will look like this: 1,2,3,5,4,6,7,8. Only 1 bribe is
// required. Print 1.
// q = [4,1,2,3]
// Person 4 had to bribe 3 people to get to the current position. Print Too chaotic.
// Function Description:
// Complete the function minimumBribes in the editor below.
// minimumBribes has the following parameter(s):
// int q[n]: the positions of the people after all bribes
// Returns
// No value is returned. Print the minimum number of bribes necessary or Too chaotic if someone has
// bribed more than 2 people.
// Input Format
// The first line contains an integer t, the number of test cases.
// Each of the next t pairs of lines are as follows:
// - The first line contains an integer t, the number of people in the queue
// - The second line has n space-separated integers describing the final state of the queue.
// Constraints
// 1 <= t <= 10
// 1 <= n <= 10^5
// Sample Input
// 2
// 5
// 2 1 5 3 4
// 5
// 2 5 1 3 4
// Sample Output
// 3
// Too chaotic
// Explanation
// Test Case 1
// The initial state:
// pic1
// After person 5 moves one position ahead by bribing person 4:
// pic2
// Now person 5 moves another position ahead by bribing person 3:
// pic3
// And person 2 moves one position ahead by bribing person 1:
// pic4
// So the final state is 2,1,5,3,4 after three bribing operations.
// Test Case 2
// No person can bribe more than two people, so its not possible to achieve the input state.

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
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

func minimumBribes(q []int32) {
	// Write your code here

	// declare const variable to hold magic number
	const (
		maxBribeEachPerson = 2
	)

	// initialize variable to hold the sum of bribes
	var numOfBribes = int32(0)

	// loop through the array backwards to:
	// outer loop: check if the difference of the value and the position in queue of the element is greater than the maximum bribe allowed
	// inner loop: by the value of the element, checks what could its 'most left' position. From there, checks if any of the elements is greater than itself,
	// if true, it means that a bribe occurred, so we increment the numOfBribes counter
	for i := int32(len(q)) - 1; i >= 0; i-- {
		if q[i]-int32(i+1) > maxBribeEachPerson {
			fmt.Println("Too chaotic")
			return
		}

		for j := max(0, (q[i]-1)-maxBribeEachPerson); j < i; j++ {
			if q[j] > q[i] {
				numOfBribes++
			}

		}
	}
	fmt.Println(numOfBribes)

}

// this function acts as a reducer for the amount of loops on the array, given that the 'most left' position of that element
// is (q[i] - 1) - maxBribeEachPerson -------> the -1 is to get the index by the value. Because the values start at '1'
func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
