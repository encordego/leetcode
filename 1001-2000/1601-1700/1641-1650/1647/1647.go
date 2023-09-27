package main

// 1647. Minimum Deletions to Make Character Frequencies Unique
//
// https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique
func minDeletions(s string) int {
	m := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		m[rune(s[i])]++
	}
	res := 0
	used := make(map[int]struct{})
	for _, freq := range m {
		_, ok := used[freq]
		for freq > 0 && ok {
			freq -= 1
			res += 1
			_, ok = used[freq]
		}
		used[freq] = struct{}{}
	}
	return res
}
