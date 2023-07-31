// Comparators are used to compare two objects. In this challenge, you'll create a comparator and use it to sort an array.
// The Player class is provided in the editor below. It has two fields:
// 1. name: a string.
// 2. score: an integer.
// Given an array of n Player objects, write a comparator that sorts them in order of decreasing score.
// If 2 or more players have the same score, sort those players alphabetically ascending by name.
// To do this, you must create a Checker class that implements the Comparator interface, then
// write an int compare(Player a, Player b) method implementing the Comparator.compare(T o1, T o2) method.
// In short, when sorting in ascending order, a comparator function returns -1 if a < b, 0 if a = b, and 1 if a > b.
// Declare a Checker class that implements the comparator method as described.
// It should sort first descending by score, then ascending by name.
// The code stub reads the input, creates a list of Player objects, uses your method to sort the data,
// and prints it out properly.

// Input Format
// Input from stdin is handled by the locked stub code in the Solution class.
// The first line contains an integer, n, denoting the number of players.
// Each of the n subsequent lines contains a player's respective name and score, a string and an integer.

// Constraints
// 0 <= score <= 1000
// 2 players can have the same name.
// Player names consist of lowercase English letters.

// Output Format
// You are not responsible for printing any output to stdout.
// The locked stub code in Solution will create a Checker object, use it to sort the Player array,
// and print each sorted element.

// Sample Input
// 5
// amy 100
// david 100
// heraldo 50
// aakansha 75
// aleksa 150

// Sample Output
// aleksa 150
// amy 100
// david 100
// aakansha 75
// heraldo 50

// Explanation
// As you can see, the players are first sorted by decreasing score and then sorted alphabetically by name.

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Read input
	var n int
	// create checker instance
	checker := Checker{}

	// scan the number of players -- first input and assign the value to the pointer address of 'n'
	fmt.Scan(&n)

	// Create a list of Player objects
	players := make([]Player, n)
	for i := 0; i < n; i++ {
		var name string
		var score int
		fmt.Scan(&name, &score)
		players[i] = Player{name, score}
	}
	// Sort the list using the Checker comparator

	sort.Slice(players, func(i, j int) bool {
		return checker.compare(players[i], players[j]) == -1
	})

	// Print the sorted list

	for _, player := range players {
		fmt.Println(player.name, player.score)
	}
}

type Player struct {
	name  string
	score int
}

type Checker struct{}

// method for the checker class, to compare two players
func (c Checker) compare(a, b Player) int {

	// if the score of both players are equal, need to sort by name asc
	if a.score == b.score {
		if a.name < b.name {
			return -1
		} else if a.name > b.name {
			return 1
		}
		return 0
	}

	// if the scores are not equal, sort by score desc
	if a.score > b.score {
		return -1
	}
	return 1
}
