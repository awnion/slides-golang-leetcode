package main

import "testing"

func TestLongestIncreasingSubsequence(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "single element",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "example from slides",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18, 2},
			expected: 4,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			name:     "leetcode example 3 all equal",
			nums:     []int{7, 7, 7, 7, 7},
			expected: 1,
		},
		{
			name:     "strictly increasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "strictly decreasing",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "two elements increasing",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "two elements decreasing",
			nums:     []int{2, 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := longestIncreasingSubsequence(tc.nums)
			if actual != tc.expected {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}
