// There is a new mobile game that starts with consecutively numbered clouds. Some of the clouds are thunderheads and
// others are cumulus. The player can jump on any cumulus cloud having a number that is equal to the number of the
// current cloud plus 1 or 2. The player must avoid the thunderheads. Determine the minimum number of jumps it will
// take to jump from the starting postion to the last cloud. It is always possible to win the game.
// For each game, you will get an array of clouds numbered 0 if they are safe or 1 if they must be avoided.
// Example
// c = [0,1,0,0,0,1,0]
// Index the array from 0...6. The number on each cloud is its index in the list so the player must avoid the clouds
// at indices 1 and 5. They could follow these two paths: 0 -> 2 -> 4 -> 6 or 0 -> 2 -> 3 -> 4 -> 6. The first path takes 3 jumps while the second takes 4.
// Return 3.
// Function Description
// 	Complete the jumpingOnClouds function in the editor below.
// 	jumpingOnClouds has the following parameter(s):
// 		int c[n]: an array of binary integers
// Returns
// 	int: the minimum number of jumps required

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
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

func jumpingOnClouds(c []int32) int32 {
	jumps := int32(0)
	// do not use magic numbers, always declare constants to represent numbers that are constant on the code
	const (
		cumulusCloud = 0
		thunderCloud = 1
		maxJumpSize  = 2
		minJumpSize  = 1
	)

	for i := 0; i < len(c)-1; {

		// simplify logic by putting 2 conditions on the if statement, instead of multiple if's inside each other
		// the conditional is only important to decide which index we go in the next loop
		// the jump incrementation stays outside that conditional, because there will always be a 1 jump incrementation
		if i+2 < len(c) && c[i+2] == cumulusCloud {
			i += maxJumpSize
		} else {
			i += minJumpSize
		}
		jumps++
	}
	return jumps
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

	cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c)

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
