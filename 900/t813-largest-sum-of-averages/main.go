package main

import "fmt"

// https://leetcode.cn/problems/largest-sum-of-averages/

func main() {
	nums := []int{9,1,2,3,9}
	k := 3

	fmt.Printf("%v", largestSumOfAverages(nums, k))
}

// 动态规划：
// dp[i][k]: 表示将 nums 中的前 i 个数，下标是 [0, i-1] 分成 k 个子数组的最大平均值和
// dp[i][k] = MAX(dp[j][k-1]+avg[j][i], dp[i][k])
// 这里是将 j 在 [0, i-1] 之间，将前 j 个数 [0, j-1] 分成 k-1 组, 加上 最后一组的平均数
// avg[j][i] 是从 [j, i-1] 范围内的平均数

func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	if n * k == 0 {
		return 0
	}
	// 前缀和
	sum := make([]float64, n+1)
	// 初始值
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + float64(nums[i])
	}
	// dp[n][k] 的二维数组
	dp := make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]float64, k+1)
	}
	// dp 值的初始化
	for i := 1; i <= n; i++ {
		// 前 i 个数分成一组，就是前 i 个数的平均值
		dp[i][1] = sum[i]/float64(i)
	}

	// 枚举分组
	for j := 2; j <= k; j++ {
		// 枚举数组元素
		for i := j; i <= n; i++ {
			// 取最佳分组位置和 这里需要从 [0,j-1] 前的所有位置都需要进行遍历， 取其中的最大值
			for x := j - 1; x < i; x++ {
				dp[i][j] = MAX(dp[i][j], dp[x][j - 1] + (sum[i] - sum[x]) / float64(i - x))
			}
		}
	}
	return dp[n][k]
}

func MAX(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}