package main

// Cherry Pickup II
// https://leetcode.com/problems/cherry-pickup-ii/
//
// Two robots start at top-left and top-right of a rows×cols grid.
// Both move down one row per step, each can shift ±1 column.
// Maximize total cherries collected; shared cell counts once.
func cherryPickupTwo(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	next, curr := make([][]int, cols), make([][]int, cols)
	for c := range cols {
		next[c] = make([]int, cols)
		curr[c] = make([]int, cols)
	}

	for c1 := range cols {
		for c2 := range cols {
			next[c1][c2] = grid[rows-1][c1]
			if c1 != c2 {
				next[c1][c2] += grid[rows-1][c2]
			}
		}
	}

	for c1 := range cols {
		for c2 := range cols {
			next[c1][c2] = grid[rows-1][c1]
			if c1 != c2 {
				next[c1][c2] += grid[rows-1][c2]
			}
		}
	}

	for r := rows - 2; r >= 0; r-- {
		for c1 := range cols {
			for c2 := range cols {
				best := 0
				for nc1 := c1 - 1; nc1 <= c1+1; nc1++ {
					for nc2 := c2 - 1; nc2 <= c2+1; nc2++ {
						if nc1 >= 0 && nc1 < cols &&
							nc2 >= 0 && nc2 < cols {
							best = max(best, next[nc1][nc2])
						}
					}
				}
				curr[c1][c2] = best + grid[r][c1]
				if c1 != c2 {
					curr[c1][c2] += grid[r][c2]
				}
			}
		}
		next, curr = curr, next
	}
	return next[0][cols-1]
}

// O(rows * cols^2) time, O(rows * cols^2) space
func cherryPickupTwoNaive(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][][]int, rows)
	for r := range rows {
		dp[r] = make([][]int, cols)
		for c1 := range cols {
			dp[r][c1] = make([]int, cols)
		}
	}

	for c1 := range cols {
		for c2 := range cols {
			dp[rows-1][c1][c2] = grid[rows-1][c1]
			if c1 != c2 {
				dp[rows-1][c1][c2] += grid[rows-1][c2]
			}
		}
	}

	for r := rows - 2; r >= 0; r-- {
		for c1 := range cols {
			for c2 := range cols {
				best := 0
				for nc1 := c1 - 1; nc1 <= c1+1; nc1++ {
					for nc2 := c2 - 1; nc2 <= c2+1; nc2++ {
						if nc1 >= 0 && nc1 < cols &&
							nc2 >= 0 && nc2 < cols {
							best = max(best, dp[r+1][nc1][nc2])
						}
					}
				}
				dp[r][c1][c2] = best + grid[r][c1]
				if c1 != c2 {
					dp[r][c1][c2] += grid[r][c2]
				}
			}
		}
	}

	return dp[0][0][cols-1]
}
