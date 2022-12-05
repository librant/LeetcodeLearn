package main

import "fmt"

// https://leetcode.cn/problems/palindromic-substrings/?favorite=2cktkvj

func main() {
	s := "abc"
	fmt.Printf("%v\n", countSubstrings(s))
}

// 中心拓展
// 奇数和偶数
// 长度为 n 的字符串会生成 2n-1 组回文中心 [li, ri]
// li = [i/2]
// ri = li + (i mod 2)
// 从 [0, 2n-2] 进行遍历 i，可以得到所有的回文中心

func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2 * n - 1; i++ {
		l, r := i / 2, i / 2 + i % 2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}