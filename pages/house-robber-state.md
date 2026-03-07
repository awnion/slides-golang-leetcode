# House Robber: state and transition

`dp[i]` = best answer considering houses `0..i`

Recurrence:

```text
dp[i] = max(
  dp[i - 1],           // skip house i
  dp[i - 2] + nums[i]  // rob house i
)
```

Base cases:

- `dp[0] = nums[0]`
- `dp[1] = max(nums[0], nums[1])`

Optimization:

- We only need the previous two states
- Space goes from `O(n)` to `O(1)`
