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
		next[c], curr[c] = make([]int, cols), make([]int, cols)
	}

	for c1 := range cols {
		for c2 := c1; c2 < cols; c2++ {
			v := grid[rows-1][c1]
			if c1 != c2 {
				v += grid[rows-1][c2]
			}
			next[c1][c2], next[c2][c1] = v, v
		}
	}

	for r := rows - 2; r >= 0; r-- {
		for c1 := range cols {
			for c2 := c1; c2 < cols; c2++ {
				best := 0
				for nc1 := max(0, c1-1); nc1 <= min(cols-1, c1+1); nc1++ {
					for nc2 := max(0, c2-1); nc2 <= min(cols-1, c2+1); nc2++ {
						best = max(best, next[nc1][nc2])
					}
				}
				v := grid[r][c1]
				if c1 != c2 {
					v += grid[r][c2]
				}
				v += best
				curr[c1][c2], curr[c2][c1] = v, v
			}
		}
		next, curr = curr, next
	}

	return next[0][cols-1]
}
