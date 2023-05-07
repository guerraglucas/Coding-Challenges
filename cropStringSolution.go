package solution

import "strings"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func SolutionCropString(message string, K int) string {
	// Implement your solution here
	if len(message) <= K {
		return message
	}

	lastSpaceIndex := strings.LastIndex(message[:K+1], " ")
	if lastSpaceIndex == -1 {
		return ""
	}
	return message[:lastSpaceIndex]
}

// TODO
// check if the string lenght is bigger than K
// if message[K] is space, will return message[:K], which K index is not included in the result
// so we need to include message[K + 1] on the check, because if the word finishes on K, we need to include it
// if there is no space, return empty string
// if there is a space, crop the string including the index of that space
