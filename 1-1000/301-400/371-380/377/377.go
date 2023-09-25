package main

var n int
var dp map[int]int

// 377. Combination Sum IV
//
// https://leetcode.com/problems/combination-sum-iv
func combinationSum4(nums []int, target int) int {
	n = len(nums)
	dp = make(map[int]int, 0)
	return solve(nums, target)
}

func solve(nums []int, sum int) int {
	if val, ok := dp[sum]; ok {
		return val
	}
	if sum == 0 {
		return 1
	}
	ans := 0
	for i := 0; i < n; i++ {
		if sum >= nums[i] {
			ans += solve(nums, sum-nums[i])
		}
	}

	dp[sum] = ans
	return ans
}
