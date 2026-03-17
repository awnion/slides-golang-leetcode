package main

import "testing"

func TestRob(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "single house",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "two houses",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "example from slides 1",
			nums:     []int{2, 7, 9, 3, 1, 5, 8, 2, 4, 6},
			expected: 26,
		},
		{
			name:     "example from slides 2",
			nums:     []int{10, 1, 1, 10, 1, 1, 10, 1, 1, 10},
			expected: 40,
		},
		{
			name:     "leetcode example 1",
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{2, 7, 9, 3, 1},
			expected: 12,
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "increasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := rob(tc.nums)
			if actual != tc.expected {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}
