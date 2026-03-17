package main

// Cherry Pickup
// https://leetcode.com/problems/cherry-pickup/
//
// n×n grid with cherries, empty cells and thorns. One robot goes (0,0)→(n-1,n-1)
// and back; cherries collected on the way there disappear. Maximize total cherries.
// Modelled as two robots walking simultaneously top-left→bottom-right,
// state: dp[r1][r2], column derived from step k: c = k - r.
func cherryPickup(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -int(1e9)
		}
	}

	dp[0][0] = grid[0][0]

	for k := 1; k < 2*n-1; k++ {
		nextDp := make([][]int, n)
		for i := range nextDp {
			nextDp[i] = make([]int, n)
			for j := range nextDp[i] {
				nextDp[i][j] = -1e9
			}
		}

		for r1 := range n {
			for r2 := range n {
				c1, c2 := k-r1, k-r2

				if c1 < 0 || c1 >= n || c2 < 0 || c2 >= n ||
					grid[r1][c1] == -1 || grid[r2][c2] == -1 {
					continue
				}

				cherries := grid[r1][c1]
				if r1 != r2 {
					cherries += grid[r2][c2]
				}

				prevMax := -int(1e9)
				for _, pr1 := range []int{r1, r1 - 1} {
					for _, pr2 := range []int{r2, r2 - 1} {
						if pr1 >= 0 && pr2 >= 0 && dp[pr1][pr2] >= 0 {
							prevMax = max(dp[pr1][pr2], prevMax)
						}
					}
				}

				if prevMax != -1e9 {
					nextDp[r1][r2] = int(prevMax) + cherries
				}
			}
		}
		dp = nextDp
	}

	res := dp[n-1][n-1]
	if res < 0 {
		return 0
	}
	return res
}
