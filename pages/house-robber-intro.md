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
