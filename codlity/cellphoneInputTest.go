package solution

import "strings"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func SolutionCellphoneInput(S string) string {
	// Implement your solution here
	var result string
	capsLock := false
	for _, char := range S {
		if string(char) == "C" {
			capsLock = !capsLock
		} else if string(char) == "B" {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			if capsLock {
				result += strings.ToUpper(string(char))
			} else {
				result += string(char)
			}
		}
	}
	return result

}

// TODO
// iterate the string and check if it is "C" or "B"
// if it is "C", change the caps lock status
// if it is "B", remove the last character
// if it is other character, append it to the result string
