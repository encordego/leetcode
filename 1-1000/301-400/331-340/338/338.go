package main

// 338. Counting Bits
//
// https://leetcode.com/problems/counting-bits
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			ans[i] = ans[i/2]
		} else {
			ans[i] = ans[i/2] + 1
		}
	}
	return ans
}
