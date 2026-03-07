# LIS: state and transition

`dp[i]` = length of the longest increasing subsequence that ends at `i`

Transition:

```text
dp[i] = 1 + max(dp[j]) for all j < i where nums[j] < nums[i]
```

If there is no such `j`, then:

```text
dp[i] = 1
```

Key idea:

- We are not asking for the best subsequence anywhere
- We are asking for the best subsequence that must end at a specific position

Final answer:

```text
max(dp[i]) for all i
```
