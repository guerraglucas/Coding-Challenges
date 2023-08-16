// HackerLand National Bank has a simple policy for warning clients about possible fraudulent account activity.
// If the amount spent by a client on a particular day is greater than or equal to 2x the client's median spending
// for a trailing number of days,
// they send the client a notification about potential fraud. The bank doesn't send the client any notifications
// until they have at least that trailing number of prior days' transaction data.
// Given the number of trailing days d and a client's total daily expenditures for a period of n days,
// find and print the number of times the client will receive a notification over all n days.

// Example
// expenditure = [10, 20, 30, 40, 50]
// d = 3
// On the first three days, they just collect spending data. At day 4, we have trailing expenditures of [10, 20, 30].
// The median is 20 and the day's expenditure is 40. Because 40 >= 2 * 20, there will be a notice.
// The next day, our trailing expenditures are [20, 30, 40] and the expenditures are 50.
// This is less than 2 * 30 so no notice will be sent. Over the period, there was one notice sent.

// Note: The median of a list of numbers can be found by arranging all the numbers from smallest to greatest.
// If there is an odd number of numbers, the middle one is picked. If there is an even number of numbers,
// median is then defined to be the average of the two middle values.

// Function Description
// Complete the function activityNotifications in the editor below.
// It must return an integer representing the number of client notifications.
// activityNotifications has the following parameter(s):
// expenditure: an array of integers representing daily expenditures
// d: an integer, the lookback days for median spending

// Returns
// int: the number of notices sent

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'activityNotifications' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY expenditure
 *  2. INTEGER d
 */
func activityNotifications(expenditure []int32, d int32) int32 {
	// Write your code here
	// check the notification case for the first d days
	countOfNotification := int32(0)

	medianFirstDDays := median(expenditure[:d])
	if expenditure[d] >= int32(2*medianFirstDDays) {
		countOfNotification++
	}
	// create a sliding window
	left := int32(1)
	right := left + d
	for right < int32(len(expenditure)) {
		todayExp := expenditure[right]
		fmt.Println(median(expenditure[left:right]))
		if todayExp >= int32(median(expenditure[left:right])*2) {
			countOfNotification++
		}
		left++
		right++
	}
	return countOfNotification
}

func median(arr []int32) float64 {
	newArr := make([]int32, len(arr))
	copy(newArr, arr)
	arrIsOdd := len(newArr)%2 != 0
	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i] < newArr[j]
	})

	if arrIsOdd {
		return float64(newArr[len(newArr)/2])
	} else {
		middleLeft := newArr[len(newArr)/2-1]
		middleRight := newArr[len(newArr)/2]
		return float64(middleLeft+middleRight) / 2
	}
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

	expenditureTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
