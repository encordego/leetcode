package main

// 1359. Count All Valid Pickup and Delivery Options
//
// https://leetcode.com/problems/count-all-valid-pickup-and-delivery-options
func countOrders(n int) int {
	mod := 1000000007
	ans := 1
	for i := 2; i <= n; i++ {
		ans = (i * (2*i - 1) * ans) % mod
	}
	return ans
}
