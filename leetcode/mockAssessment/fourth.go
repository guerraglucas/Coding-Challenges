// You are given an m x n grid where each cell can have one of three values:

// 0 representing an empty cell,
// 1 representing a fresh orange, or
// 2 representing a rotten orange.
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

// Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

// Example 1:

// Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
// Output: 4

// Example 2:

// Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
// Output: -1
// Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

// Example 3:

// Input: grid = [[0,2]]
// Output: 0
// Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.

package mockAssessment

type Point struct {
	row, col int
}

func OrangesRotting(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	freshCount := 0 // Count of fresh oranges
	queue := []Point{}

	// Count fresh oranges and initialize the queue with rotten oranges
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				freshCount++
			} else if grid[i][j] == 2 {
				queue = append(queue, Point{i, j})
			}
		}
	}

	if freshCount == 0 {
		return 0 // No fresh oranges initially
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				newRow, newCol := point.row+dir[0], point.col+dir[1]
				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2 // Set the fresh orange as rotten
					freshCount--
					queue = append(queue, Point{newRow, newCol})
				}
			}
		}
		minutes++
	}

	if freshCount == 0 {
		return minutes - 1 // All fresh oranges have become rotten
	}
	return -1 // Impossible to rot all fresh oranges
}
