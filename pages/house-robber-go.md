# House Robber in Go

```go
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	prev := nums[0]
	cur := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		prev, cur = cur, max(cur, prev+nums[i])
	}

	return cur
}
```
