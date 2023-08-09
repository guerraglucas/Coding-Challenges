// a fibonacci with modulo sequence is defined for a given modulo, mod as follows:
// fib[x] = (fib[x-1] + fib[x-2]) % mod, x>2
// fib[1] = 1, fib[2] = 2

// [1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...] are the fiurst few termos of the
// fibonacci sequence. In this task , you are required to write a function
// that run a server that generates a fibonacci sequence

// it accepts boolean requests and returns the next number in the sequence.
// it should have a rate limiter which should make the server send every response no earilier than
// 10 milliseconds after the previous one.
// The server should be created with two arguments - a boolean channel that accepts requests and a result
// int channel trough which every result will be returned. The main funcion will also accept two
// arguments: the number of results that will be skipped from the beginning, and the number of results
// that will be printed to the console.

// note: since the numbers can be large, the server must return the number modulo 10^9.
// for example, for the 50th fibonacci number - 12586269025, the server should return 586269025.

//for example, using the arguments 2 and 5 results in [3, 5, 8, 13, 21]

// sample imput
// STDIN    Function
// -----    --------
// 0        -> skip = 0
// 6 	    -> total = 6

// sample output
// 1
// 2
// 3
// 5
// 8
// 13

// explanation
// 0 means no elements will be skipped, so the sequence starts with the first element, 1.
// the 6 elements in the sequence are 1, 2, 3, 5, 8, 13. none of these is large enouth for the modulo to change

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'ModuloFibonacciSequence' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 */

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	fib := []int{1, 2}

	// Define a time for rate limiting
	rateLimit := time.Tick(10 * time.Millisecond)

	sendResult := func(value int) {
		resultChan <- value
	}

	for i := 2; i > 0; i-- {
		<-requestChan
		sendResult(fib[2-i])
		<-rateLimit
	}

	for {
		<-requestChan
		new := (fib[len(fib)-1] + fib[len(fib)-2]) % 1000000000
		fib = append(fib, new)
		sendResult(new)
		<-rateLimit
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	skip := int32(skipTemp)

	totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	total := int32(totalTemp)

	resultChan := make(chan int)
	requestChan := make(chan bool)
	go ModuloFibonacciSequence(requestChan, resultChan)
	for i := int32(0); i < skip+total; i++ {
		start := time.Now().UnixNano()
		requestChan <- true
		new := <-resultChan
		if i < skip {
			continue
		}
		end := time.Now().UnixNano()
		timeDiff := (end - start) / 1000000
		if timeDiff < 3 {
			fmt.Println("Rate is too high")
			break
		}
		fmt.Println(new)
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
