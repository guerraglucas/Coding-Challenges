// Harold is a kidnapper who wrote a ransom note, but now he is worried it will be traced back to him through his handwriting.
// He found a magazine and wants to know if he can cut out whole words from it and use them to create an untraceable replica of his ransom note.
// The words in his note are case-sensitive and he must use only whole words available in the magazine.
// He cannot use substrings or concatenation to create the words he needs.
// Given the words in the magazine and the words in the ransom note, print Yes if he can replicate his ransom note exactly using whole words from the magazine; otherwise, print No.
// For example, the note is "Attack at dawn".
// The magazine contains only "attack at dawn". The magazine has all the right words, but there's a case mismatch. The answer is No.
// Function Description
// Complete the checkMagazine function in the editor below. It must print Yes if the note can be formed using the magazine, or No.
// checkMagazine has the following parameters:
// magazine: an array of strings, each a word in the magazine
// note: an array of strings, each a word in the ransom note
// Input Format
// The first line contains two space-separated integers, m and n, the numbers of words in the magazine and the note..
// The second line contains m space-separated strings, each magazine[i].
// The third line contains n space-separated strings, each note[i].
// Constraints
// 1 <= m, n <= 30000
// 1 <= |magazine[i]|, |note[i]| <= 5
// Each word consists of English alphabetic letters (i.e., a to z and A to Z).
// Output Format
// Print Yes if he can use the magazine to create an untraceable replica of his ransom note. Otherwise, print No.
// Sample Input 0
// 6 4
// give me one grand today night
// give one grand today
// Sample Output 0
// Yes
// Sample Input 1
// 6 5
// two times three is not four
// two times two is four
// Sample Output 1
// No
// Explanation 1
// 'two' only occurs once in the magazine.
// Sample Input 2
// 7 4
// ive got a lovely bunch of coconuts
// ive got some coconuts
// Sample Output 2
// No

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
 * Complete the 'checkMagazine' function below.
 *
 * The function accepts following parameters:
 *  1. STRING_ARRAY magazine
 *  2. STRING_ARRAY note
 */

func checkMagazine(magazine []string, note []string) {
	// Write your code here
	mapOfMagazine := make(map[string]int)
	for _, word := range magazine {
		mapOfMagazine[word]++
	}
	for _, word := range note {
		if _, ok := mapOfMagazine[word]; ok {
			if mapOfMagazine[word] > 0 {
				mapOfMagazine[word]--
				continue

			} else {
				fmt.Println("No")
				return
			}
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	mTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
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
