package main

import "testing"

func TestCherryPickup(t *testing.T) {
	testCases := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "LeetCode Example 1",
			grid: [][]int{
				{0, 1, -1},
				{1, 0, -1},
				{1, 1, 1},
			},
			expected: 5,
		},
		{
			name: "Smallest grid",
			grid: [][]int{
				{1},
			},
			expected: 1,
		},
		{
			name: "2x2 all cherries",
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			expected: 4,
		},
		{
			name: "3x3 all cherries",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 8,
		},
		{
			name: "Grid with walls - only one path",
			grid: [][]int{
				{0, 1, 1},
				{-1, -1, 1},
				{1, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Grid with walls - only one path 2",
			grid: [][]int{
				{0, -1, 1},
				{1, -1, 1},
				{1, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Grid where crossing is optimal",
			grid: [][]int{
				{1, 1, 1, 1, 0},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 1, 1},
			},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := cherryPickup(tc.grid)
			if actual != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, actual)
			}
		})
	}
}
