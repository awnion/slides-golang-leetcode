package main

import "testing"

func TestCherryPickupTwo(t *testing.T) {
	testCases := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "1x1 grid",
			grid: [][]int{
				{5},
			},
			expected: 5,
		},
		{
			name: "1x5 grid (linear)",
			grid: [][]int{
				{1, 2, 3, 4, 5},
			},
			expected: 6,
		},
		{
			name: "5x1 grid (linear)",
			grid: [][]int{
				{1}, {2}, {3}, {4}, {5},
			},
			expected: 15,
		},
		{
			name: "2x2 grid all cherries",
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			expected: 4,
		},
		{
			name: "3x3 LeetCode style",
			grid: [][]int{
				{0, 1, -1},
				{1, 0, -1},
				{1, 1, 1},
			},
			expected: 2,
		},
		{
			name: "3x3 all 1s",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 6,
		},
		{
			name:     "5x10 all 1s",
			grid:     makeAllOnes(5, 10),
			expected: 10,
		},
		{
			name: "5x10 with walls",
			grid: [][]int{
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := cherryPickupTwo(tc.grid)
			if actual != tc.expected {
				t.Errorf("%s: expected %d, but got %d", tc.name, tc.expected, actual)
			}
		})
	}
}

func makeAllOnes(r, c int) [][]int {
	grid := make([][]int, r)
	for i := range grid {
		grid[i] = make([]int, c)
		for j := range grid[i] {
			grid[i][j] = 1
		}
	}
	return grid
}
