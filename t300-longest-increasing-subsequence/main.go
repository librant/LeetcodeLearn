package main

import "fmt"

// https://leetcode.cn/problems/longest-increasing-subsequence/?favorite=2cktkvj

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

// 动态规划：
// dp[i]: 表示前 i 个数中，严格递增子数组的长度
// dp[j]: 表示前 j 个数中，严格递增子数组的长度 [0, i-1]
// 状态转移方程：
// 当 nums[i] > nums[j], 则表示可以接在 dp[j] 后面, 此时， dp[i] = MAX(dp[j]) + 1
// 当 nums[i] <= nums[j], 则表示不可以接在 dp[j] 后面, 此时 dp[i] = dp[j]
// 注意 {nums}[i]nums[i] 必须被选取
// 这里的最大值，不一定是 dp[n-1]， 因为这个 dp[i] 是必须选择 num[i] 的值，所以要取 dp[i] 中最大的值

func main() {
	nums1 := []int{10,9,2,5,3,7,101,18}
	fmt.Printf("%v\n", lengthOfLIS(nums1))

	nums2 := []int{4,10,4,3,8,9}
	fmt.Printf("%v\n", lengthOfLIS(nums2))

}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	// 初始化
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 这里知道最大的 dp[j], 先找出所有 dp[j] 中的最大值
				dp[i] = MAX(dp[i], dp[j]+1)
				res = MAX(dp[i], res)
			}
		}
	}
	return res
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}