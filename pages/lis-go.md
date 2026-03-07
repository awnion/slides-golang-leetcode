# Longest Increasing Subsequence in Go

```go
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	best := 1

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		best = max(best, dp[i])
	}

	return best
}
```
