package main

// House Robber
// https://leetcode.com/problems/house-robber/
//
// Row of houses with gold. Can't rob two adjacent houses. Maximize total loot.
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	prev, cur := nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		prev, cur = cur, max(cur, prev+nums[i])
	}

	return cur
}
