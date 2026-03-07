---
theme: default
title: Dynamic Programming Through LeetCode
titleTemplate: '%s'
info: |
  A Slidev deck about dynamic programming patterns in interview-style problems.
class: text-left
drawings:
  persist: false
transition: slide-left
mdc: true
fonts:
  sans: IBM Plex Sans
  mono: afio
---

<div class="hero-grid">
  <div>
    <div class="hero-eyebrow">Job Interviews</div>
    <h1 class="hero-title">Dynamic Programming</h1>
    <div class="hero-subtitle">Selected LeetCode Problems in Go</div>
    <div class="hero-copy">
      A structured walkthrough of interview-style DP cases:
    </div>
    <div class="hero-points">
      <div><span class="hero-point-label">Pattern</span> state -> transition -> optimization</div>
      <div><span class="hero-point-label">Scope</span> more than one problem, from lighter to heavier cases</div>
      <div><span class="hero-point-label">Context</span> what actually matters at the coding interview stage</div>
    </div>
    <div class="hero-author">Sergei B. at AckiNacki</div>
  </div>

  <div class="hero-visual">
    <div class="hero-card">
      <div class="hero-card-top">Interview Loop</div>
      <img class="hero-gopher" src="/gopher-running.svg" alt="Go gopher running hard toward the next interview round" />
      <div class="hero-caption">LeetCode round: survive the recurrence, earn the next conversation.</div>
    </div>
  </div>
</div>

<div class="hero-credit">Go gopher artwork by Renee French via go.dev, CC BY 4.0.</div>

---
layout: center
---

# About Me

`Your Name`

`Role / Team / Company`

`What I build`

`Why I care about interviews / algorithms`

<div class="pt-8 text-sm text-slate-500">
Replace this slide with your actual intro.
</div>

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
  <div class="text-sm uppercase tracking-[0.2em] text-amber-600 font-semibold">Reality Check</div>
  <div class="pt-3 text-lg">
    LeetCode is <span class="accent">not the job</span>, but it is still part of the hiring pipeline.
  </div>
  <div class="pt-4 tiny text-slate-600">
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
  <div class="chip font-bold border-teal-600 text-teal-800">Dynamic Programming</div>
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
  <div class="text-sm uppercase tracking-[0.2em] text-amber-600 font-semibold">Often Confused With</div>
  <div class="pt-3">
    <div><span class="accent">Recursion</span>: a way to express a solution</div>
    <div class="pt-2"><span class="accent">Memoization</span>: a tool that often implements DP</div>
    <div class="pt-2"><span class="accent">Greedy</span>: making the best local choice, which often fails for DP problems</div>
    <div class="pt-2"><span class="accent">Backtracking</span>: exploring choices, usually without reusing computed results</div>
  </div>
</div>

<div class="pt-6 tiny text-slate-600">
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

	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		cur := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = cur
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

<div class="pt-4 tiny text-slate-600">Time: O(n) • Space: O(1)</div>

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

<div class="pt-4 tiny text-slate-600">Time: O(n²) • Space: O(n)</div>

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

<div class="pt-8 text-slate-600">
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

<div class="pt-6 text-slate-600">
`your.name@example.com` • `github.com/your-handle`
</div>

<style>
:root {
  --slidev-theme-primary: #0f766e;
  --brand-deep: #0b132b;
  --brand-teal: #0f766e;
  --brand-amber: #f59e0b;
  --brand-slate: #334155;
}

.slidev-layout {
  background:
    radial-gradient(circle at top right, rgba(15, 118, 110, 0.16), transparent 28%),
    radial-gradient(circle at bottom left, rgba(245, 158, 11, 0.14), transparent 24%),
    linear-gradient(180deg, #fcfcfb 0%, #f8fafc 100%);
  color: var(--brand-deep);
}

.hero-grid {
  display: grid;
  grid-template-columns: minmax(0, 1.2fr) minmax(280px, 0.8fr);
  gap: 2rem;
  align-items: center;
  min-height: 74vh;
}

.hero-eyebrow {
  text-transform: uppercase;
  letter-spacing: 0.28em;
  font-size: 0.76rem;
  font-weight: 700;
  color: var(--brand-teal);
}

.hero-title {
  margin-top: 0.8rem;
  margin-bottom: 0;
  max-width: 10ch;
  font-family: "Iowan Old Style", "Palatino Linotype", "Book Antiqua", Georgia, serif;
  font-size: 3.45rem;
  font-weight: 700;
  line-height: 0.93;
  color: #0f172a;
}

.hero-subtitle {
  margin-top: 0.7rem;
  max-width: 22rem;
  font-size: 1.12rem;
  font-weight: 700;
  line-height: 1.35;
  color: #0f766e;
}

.hero-copy {
  margin-top: 1.1rem;
  max-width: 36rem;
  font-size: 1rem;
  line-height: 1.6;
  color: var(--brand-slate);
}

.hero-points {
  margin-top: 0.9rem;
  display: grid;
  gap: 0.55rem;
  max-width: 38rem;
}

.hero-points > div {
  border-left: 3px solid rgba(15, 118, 110, 0.22);
  padding-left: 0.85rem;
  color: #1e293b;
  line-height: 1.45;
}

.hero-point-label {
  margin-right: 0.35rem;
  color: #b45309;
  font-weight: 700;
}

.hero-author {
  margin-top: 1.4rem;
  display: inline-flex;
  align-items: center;
  border: 1px solid rgba(15, 118, 110, 0.24);
  border-radius: 999px;
  padding: 0.45rem 0.9rem;
  background: linear-gradient(90deg, rgba(255, 255, 255, 0.9), rgba(240, 253, 250, 0.95));
  font-size: 0.95rem;
  font-weight: 600;
}

.hero-visual {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 0.8rem 0.4rem 0.4rem;
}

.hero-card {
  position: relative;
  width: min(100%, 360px);
  overflow: hidden;
  border: 1px solid rgba(15, 23, 42, 0.08);
  border-radius: 1.6rem;
  background:
    radial-gradient(circle at top left, rgba(56, 189, 248, 0.38), transparent 32%),
    radial-gradient(circle at bottom right, rgba(245, 158, 11, 0.28), transparent 30%),
    linear-gradient(160deg, rgba(255, 255, 255, 0.98), rgba(240, 249, 255, 0.95));
  padding: 1rem 1rem 1.15rem;
  box-shadow: 0 24px 50px rgba(15, 23, 42, 0.11);
}

.hero-card::before,
.hero-card::after {
  content: "";
  position: absolute;
  border-radius: 999px;
}

.hero-card::before {
  top: 0.8rem;
  right: 0.9rem;
  width: 0.9rem;
  height: 0.9rem;
  background: #f59e0b;
  box-shadow: -1.3rem 0 0 #14b8a6, -2.6rem 0 0 #38bdf8;
}

.hero-card::after {
  left: -2.4rem;
  bottom: -2.8rem;
  width: 7rem;
  height: 7rem;
  background: rgba(15, 118, 110, 0.1);
}

.hero-gopher {
  position: relative;
  z-index: 1;
  display: block;
  margin: 0.4rem auto 0;
  width: min(100%, 280px);
  height: auto;
  filter: drop-shadow(0 16px 30px rgba(15, 23, 42, 0.16));
}

.hero-card-top {
  position: relative;
  z-index: 1;
  display: inline-flex;
  align-items: center;
  border-radius: 999px;
  background: rgba(15, 118, 110, 0.12);
  padding: 0.35rem 0.75rem;
  font-size: 0.76rem;
  font-weight: 800;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: #0f766e;
}

.hero-caption {
  position: relative;
  z-index: 1;
  margin-top: 0.4rem;
  border-radius: 1rem;
  background: rgba(255, 255, 255, 0.72);
  padding: 0.75rem 0.9rem;
  font-size: 0.86rem;
  line-height: 1.45;
  color: var(--brand-deep);
  text-align: center;
}

.hero-credit {
  margin-top: -0.5rem;
  font-size: 0.68rem;
  color: rgba(51, 65, 85, 0.78);
}

code,
pre,
kbd,
samp,
.font-mono,
.shiki {
  font-family: afio, monospace;
}

h1, h2, h3 {
  letter-spacing: -0.02em;
}

.lead {
  font-size: 1.2rem;
  line-height: 1.6;
  color: var(--brand-slate);
}

.accent {
  color: var(--brand-teal);
  font-weight: 700;
}

.chip-row {
  display: flex;
  flex-wrap: wrap;
  gap: 0.6rem;
  margin-top: 1rem;
}

.chip {
  border: 1px solid rgba(15, 118, 110, 0.2);
  background: rgba(255, 255, 255, 0.8);
  border-radius: 999px;
  padding: 0.35rem 0.8rem;
  font-size: 0.95rem;
}

.problem-card {
  border: 1px solid rgba(15, 23, 42, 0.08);
  border-radius: 1rem;
  padding: 1rem 1.1rem;
  background: rgba(255, 255, 255, 0.85);
  box-shadow: 0 10px 30px rgba(15, 23, 42, 0.05);
}

.tiny {
  font-size: 0.92rem;
}

@media (max-width: 900px) {
  .hero-grid {
    grid-template-columns: 1fr;
    gap: 1.25rem;
    min-height: auto;
  }

  .hero-title {
    max-width: none;
    font-size: 2.7rem;
  }

  .hero-copy {
    max-width: none;
  }
}
</style>
