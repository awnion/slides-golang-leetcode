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
