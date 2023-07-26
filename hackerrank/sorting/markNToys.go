// Mark and Jane are very happy after having their first child. Their son loves toys, so Mark wants to buy some.
// There are a number of different toys lying in front of him, tagged with their prices.
// Mark has only a certain amount to spend, and he wants to maximize the number of toys he buys with this money.
// Given a list of prices and an amount to spend, what is the maximum number of toys Mark can buy?
// Note: Mark can buy only 1 of each toy.
// Example = [1, 2, 3, 4]
// k = 7
// The budget is 7 units of currency. He can buy items that cost [1, 2, 3] for 6, or [3, 4] for 7 units.
// The maximum is 3 items.

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
 * Complete the 'maximumToys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY prices
 *  2. INTEGER k
 */

func maximumToys(prices []int32, k int32) int32 {
	// Write your code here
	count := int32(0)
	sumPrices := int32(0)
	bubbleSort(&prices)

	for _, value := range prices {
		if sumPrices+value <= k {
			sumPrices += value
			count++
		} else {
			break
		}
	}

	return count

}

func bubbleSort(arr *[]int32) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
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

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	pricesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var prices []int32

	for i := 0; i < int(n); i++ {
		pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
		checkError(err)
		pricesItem := int32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}

	result := maximumToys(prices, k)

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
