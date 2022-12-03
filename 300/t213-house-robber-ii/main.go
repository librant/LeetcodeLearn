package main

import "fmt"

// https://leetcode.cn/problems/house-robber-ii/

// 次题房屋是个环行
// 可以分成两部分计算：
// 偷第一间，就不能偷最后一间 [0, n-2]
// 偷最后一间，就不能偷第一间 [1, n-1]
// dp 参考 t198 题

func main() {
	nums := []int{2,3,2}
	fmt.Printf("%v\n", rob(nums))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return MAX(nums[0], nums[1])
	}

	return MAX(dprob(nums[:n-1]), dprob(nums[1:]))
}


// 动态规划：
// dp[i] = MAX(dp[j], dp[i])
func dprob(nums []int) int {
	n := len(nums)
	max := 0
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = MAX(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = MAX(dp[i-1], dp[i-2] + nums[i])
	}
	for i := 0; i < n; i++ {
		max = MAX(dp[i], max)
	}
	return max
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}
