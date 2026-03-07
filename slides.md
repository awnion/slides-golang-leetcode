---
theme: default
title: Dynamic Programming Through LeetCode
titleTemplate: "%s"
info: |
  A Slidev deck about dynamic programming patterns in interview-style problems.
class: text-left
drawings:
  persist: false
transition: slide-left
mdc: true
highlighter: shiki
shiki:
  light: github-light
  dark: one-dark-pro
fonts:
  sans: IBM Plex Sans
  mono: afio
layout: center
---

<div class="grid grid-cols-[1fr_1fr] gap-4 items-center w-full">
  <div>
    <div class="uppercase text-base font-bold text-[var(--brand-accent)] text-4xl">LeetCode Problems on Go</div>
    <div class="mt-2 font-bold text-6xl" style="font-family: 'Iowan Old Style', 'Palatino Linotype', Georgia, serif;">Dynamic Programming</div>
    <div class="mt-12 items-center color-[var(--text-secondary)]">Sergei Blinov · AckiNacki</div>
  </div>

  <div class="flex flex-col items-center justify-center">
    <div class="relative w-full overflow-hidden border border-[var(--border-subtle)] rounded-3xl bg-[var(--bg-card)] p-3 shadow-2xl">
      <img class="block w-full h-auto rounded-xl" src="/gopher-this-is-fine.png" alt="Go gopher 'this is fine' meme" />
      <div class="mt-2 rounded-2xl bg-[var(--bg-card)] px-4 py-2 text-lg text-center text-[var(--text-secondary)]">Solving DP problems at a coding interview</div>
    </div>
  </div>
</div>

---
layout: center
---

# About Me

Sergei Blinov

- `Senior Software Engineer at AckiNacki`
- `Math enthusiast`
- `Programming language polyglot`

---
layout: two-cols
---

# Why companies still ask LeetCode

<div class="lead">
Not because the job is "reverse a binary tree" every day, but because interviews still need a fast and somewhat standardized signal.
</div>

- It tests structured problem solving under time pressure
- It exposes how candidates model tradeoffs
- It is easier to calibrate than open-ended system design for junior and mid-level roles
- It creates a common language across teams and interviewers

::right::

<div class="problem-card mt-8">
  <div class="text-sm uppercase tracking-[0.2em] text-[var(--brand-accent)] font-semibold">Reality Check</div>
  <div class="pt-3 text-lg">
    LeetCode is <span class="accent">not the job</span>, but it is still part of the hiring pipeline.
  </div>
  <div class="pt-4 tiny text-[var(--text-muted)]">
    Especially when companies need a scalable, repeatable interview process.
  </div>
</div>

---

# Common interview topics

<div class="chip-row">
  <div class="chip">Arrays & Hashing</div>
  <div class="chip">Two Pointers</div>
  <div class="chip">Binary Search</div>
  <div class="chip">Stacks & Queues</div>
  <div class="chip">Trees & Graphs</div>
  <div class="chip">Backtracking</div>
  <div class="chip">Greedy</div>
  <div class="chip">Heap / Priority Queue</div>
  <div class="chip font-bold border-[var(--go-blue)] text-[var(--brand-accent)]">Dynamic Programming</div>
</div>

<div class="pt-10 lead">
Dynamic programming tends to be the turning point because it forces you to define:
</div>

- what the subproblem is
- what information must be remembered
- how a local choice affects the future

---
layout: section
---

# 3 DP Problems

From a linear recurrence to a high-dimensional state space

---
layout: two-cols
---

# Why is it called dynamic programming?

<div class="lead">
The name is historical, not descriptive.
It does <span class="accent">not</span> mean "programming" as in writing code, and it does <span class="accent">not</span> mean "dynamic" as in changing over time.
</div>

- Richard Bellman popularized the term in the context of optimization
- In practice, DP means:
  solving a problem by combining answers to overlapping subproblems
- The real work is choosing the right state and recurrence

::right::

<div class="problem-card mt-4">
  <div class="text-sm uppercase tracking-[0.2em] text-[var(--brand-accent)] font-semibold">Often Confused With</div>
  <div class="pt-3">
    <div><span class="accent">Recursion</span>: a way to express a solution</div>
    <div class="pt-2"><span class="accent">Memoization</span>: a tool that often implements DP</div>
    <div class="pt-2"><span class="accent">Greedy</span>: making the best local choice, which often fails for DP problems</div>
    <div class="pt-2"><span class="accent">Backtracking</span>: exploring choices, usually without reusing computed results</div>
  </div>
</div>

<div class="pt-6 tiny text-[var(--text-muted)]">
DP is usually about structured reuse of computed states.
</div>

---

# Problem 1: House Robber

<div class="lead">
Easy to explain, easy to brute force, but already useful for introducing the core DP habit:
<span class="accent">at index i, what is the best answer so far?</span>
</div>

<div class="problem-card mt-6">
Given an integer array `nums`, return the maximum amount of money you can rob tonight without robbing two adjacent houses.
</div>

- Brute force asks: rob this house or skip it?
- DP asks: what is the best total ending at position `i`?
- This is a classic 1D recurrence

---

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

---

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

---

# Problem 2: Longest Increasing Subsequence

<div class="lead">
Now the state is still 1D, but the transition is no longer local.
Each position may depend on <span class="accent">many previous positions</span>.
</div>

<div class="problem-card mt-6">
Given an integer array `nums`, return the length of the longest strictly increasing subsequence.
</div>

- This is a good "medium" DP problem
- It teaches how to define "best answer ending at index i"
- It also opens the door to discussing better-than-DP solutions later

---

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

---

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

---

# Problem 3: Cherry Pickup II

<div class="lead">
This is where DP starts to feel genuinely hard: the state must represent
<span class="accent">two agents moving at the same time</span>.
</div>

<div class="problem-card mt-6">
Two robots start at the top row of a grid, one on the left and one on the right. On each step, both move to the next row and may shift by `-1`, `0`, or `+1` columns. Maximize the total cherries collected.
</div>

- A greedy approach fails quickly
- A naive DFS explodes because the branching factor is large
- The key is to encode both robot positions in the state

---

# Cherry Pickup II: state design

State:

```text
dp[row][col1][col2]
```

Meaning:

- we are on `row`
- robot A is at `col1`
- robot B is at `col2`
- value = maximum cherries collectable from this state to the bottom

Transition:

- each robot has 3 moves
- total next states per step = `3 x 3 = 9`

Reward on the current row:

```text
grid[row][col1] + grid[row][col2]
```

If `col1 == col2`, count that cell only once.

---

# Cherry Pickup II: recurrence

```text
dp[row][c1][c2] =
  cherries(row, c1, c2) +
  max(
    dp[row + 1][nc1][nc2]
    for nc1 in {c1 - 1, c1, c1 + 1}
    for nc2 in {c2 - 1, c2, c2 + 1}
  )
```

Important details:

- invalid columns must be skipped
- the last row is the base case
- top-down memoization and bottom-up DP both work

Complexity:

- States: `O(rows * cols * cols)`
- Transitions per state: constant (`9`)
- Total: `O(rows * cols²)`

---

# Cherry Pickup II in Go

```go
func cherryPickup(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][][]int, rows)
	for r := 0; r < rows; r++ {
		dp[r] = make([][]int, cols)
		for c1 := 0; c1 < cols; c1++ {
			dp[r][c1] = make([]int, cols)
		}
	}

	for c1 := 0; c1 < cols; c1++ {
		for c2 := 0; c2 < cols; c2++ {
			dp[rows-1][c1][c2] = grid[rows-1][c1]
			if c1 != c2 {
				dp[rows-1][c1][c2] += grid[rows-1][c2]
			}
		}
	}

	for r := rows - 2; r >= 0; r-- {
		for c1 := 0; c1 < cols; c1++ {
			for c2 := 0; c2 < cols; c2++ {
				best := 0
				for d1 := -1; d1 <= 1; d1++ {
					for d2 := -1; d2 <= 1; d2++ {
						nc1, nc2 := c1+d1, c2+d2
						if nc1 >= 0 && nc1 < cols && nc2 >= 0 && nc2 < cols {
							best = max(best, dp[r+1][nc1][nc2])
						}
					}
				}
				dp[r][c1][c2] = grid[r][c1]
				if c1 != c2 {
					dp[r][c1][c2] += grid[r][c2]
				}
				dp[r][c1][c2] += best
			}
		}
	}

	return dp[0][0][cols-1]
}
```

---

# What these three problems teach

1. `House Robber`: a DP state can be tiny and still powerful
2. `LIS`: the hard part is often defining the right subproblem
3. `Cherry Pickup II`: complexity jumps when the state tracks multiple moving pieces

<div class="pt-8 lead">
The recurring checklist is always the same:
</div>

- What does my state mean?
- What choices move me to the next state?
- What is the base case?
- Can I reduce memory without losing information?

---

# References

- [House Robber](https://leetcode.com/problems/house-robber/)
- [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)
- [Cherry Pickup II](https://leetcode.com/problems/cherry-pickup-ii/)
- [Slidev Documentation](https://sli.dev/)

<div class="pt-8 text-[var(--text-muted)]">
Optional ending ideas:
</div>

- "If you panic in DP, start by naming the state."
- "Interview DP is usually less about magic and more about bookkeeping."
- Add your GitHub / LinkedIn / QR code here.

---
layout: center
class: text-center
---

# Thank You

Questions?

<div class="pt-6 text-[var(--text-muted)]">
`your.name@example.com` • `github.com/your-handle`
</div>
