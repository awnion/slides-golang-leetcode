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
