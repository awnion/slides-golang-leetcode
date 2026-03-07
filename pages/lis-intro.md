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
