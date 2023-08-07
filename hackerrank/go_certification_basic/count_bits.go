// Implement a function that counts the number of set bits in the binary representation of a 32-bit integer.
// num = 126
// 126 = 1111110 has 6 bits set
// num = 128
// 128 = 10000000 has 1 bit set

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
 * Complete the 'countBits' function below.
 *
 * The function is expected to return an int32.
 * The function accepts unit32 num as parameter.
 */

func countBits(num uint32) int32 {
	binary := fmt.Sprintf("%032b", num)
	count := int32(0)
	for _, set := range binary {
		if string(set) == "1" {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	numTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	num := uint32(numTemp)

	result := countBits(num)

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
