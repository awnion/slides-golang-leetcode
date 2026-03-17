package main

// Longest Increasing Subsequence
// https://leetcode.com/problems/longest-increasing-subsequence/
//
// Find the length of the longest strictly increasing subsequence.
// Elements don't have to be adjacent.
func longestIncreasingSubsequence(nums []int) int {
	dp := make([]int, len(nums))
	globalMax := 1

	for i := range nums {
		localMax := 1
		for j := range i {
			if nums[j] < nums[i] {
				localMax = max(localMax, dp[j]+1)
			}
		}
		dp[i], globalMax = localMax, max(globalMax, localMax)
	}

	return globalMax
}
