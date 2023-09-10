package main

// 2707. Extra Characters in a String
//
// https://leetcode.com/problems/extra-characters-in-a-string/
func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i] + 1
		for _, word := range dictionary {
			l := len(word)
			if i >= l-1 && word == s[i-l+1:i+1] && dp[i-l+1] < dp[i+1] {
				dp[i+1] = dp[i-l+1]
			}
		}
	}
	return dp[n]
}
