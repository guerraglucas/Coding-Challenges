package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func SolutionTreeDistinctPath(T *Tree) int {
	// Implement your solution here

	// creates a set to keep track of the values
	set := make(map[int]bool)

	// calls the recursive function
	return distinctPath(T, set, 0)
}

func distinctPath(node *Tree, set map[int]bool, count int) int {

	// returns if node is nil
	if node == nil {
		return count
	}

	// checks if the node value is already on the set
	if _, ok := set[node.X]; ok {
		return count
	}

	// adds the node value to the set
	set[node.X] = true

	// recursively calls for each side nodes incrementing the count
	leftCount := distinctPath(node.L, set, count+1)
	rightCount := distinctPath(node.R, set, count+1)

	// removes the node value from the set
	delete(set, node.X)

	// returns the highest count
	if leftCount > rightCount {
		return leftCount
	}

	return rightCount
}
