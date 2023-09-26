// Each time Sunny and Johnny take a trip to the Ice Cream Parlor, they pool together m dollars for ice cream.
// On any given day, the parlor offers a line of n flavors. Each flavor, i, is numbered sequentially with a
// unique ID number from 1 to n and has a cost, ci, associated with it.
//
// Given the value of m and the cost of each flavor for t trips to the Ice Cream Parlor, help Sunny and Johnny
// choose two flavors such that they spend their entire pool of money (m) during each visit. ID numbers are the 1-
// based index number associated with a cost. For each trip to
// the parlor, print the ID numbers for the two types of ice cream that Sunny and Johnny purchase as two
// space-separated integers on a new line. You must print the smaller ID first and the larger ID second.
//
// Example
// cost = [2,1,3,5,6]
// m = 5
// They would purchase flavor ID's 1 and 3 for a cost of 2 + 3 = 5. Use 1 based indexing for your response.
//
// Note:
// Two ice creams having unique IDs i and j may have the same cost (i.e., ci = cj).
// There will always be a unique solution.

// Function Description
// Complete the function whatFlavors in the editor below.
// whatFlavors has the following parameter(s):
// int cost[n]: the prices for each flavor
// int money: the amount of money they have to spend
//
// Prints
// - int int: the indices of the two flavors they will purchase as two space-separated integers on a line

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
 * Complete the 'whatFlavors' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY cost
 *  2. INTEGER money
 */

func whatFlavors(cost []int32, money int32) {
	// Write your code here

	mapOfFlavors := make(map[int32]int32)
	for index, flavor := range cost {
		remaining := money - flavor

		if _, ok := mapOfFlavors[remaining]; ok {
			fmt.Println(mapOfFlavors[remaining]+1, index+1)
			break
		}

		mapOfFlavors[flavor] = int32(index)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
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
