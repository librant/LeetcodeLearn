package main

import "fmt"

// https://leetcode.cn/problems/maximal-square/

func main() {
	matrix1 := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	fmt.Printf("%v\n", maximalSquare(matrix1))

	matrix2 := [][]byte{{'0','1'},{'1','0'}}
	fmt.Printf("%v\n", maximalSquare(matrix2))
}

// 动态规划：
// dp[i][j]: 表示以 (i, j) 为右下角坐标，且只包含 1 的正方形的最大边长；
// (i, j) 如果为 '0'， 则 dp[i][j] = 0
// (i, j) 如果为 '1'， 则 dp[i][j] 的值由 dp[i-1][j], dp[i][j-1], dp[i-1][j-1] 三者中最小值加 1
// 包围 (i, j) 前的最小正方形，边长再加 1

func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	dp := make([][]int, n)
	res := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(matrix[i]))
		// 对第 j = 0 列进行初始化
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			res = MAX(res, dp[i][0])
		}
 	}
	// 对第 i = 0 行进行初始化
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			res = MAX(res, dp[0][j])
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = MIN(MIN(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				res = MAX(res, dp[i][j] * dp[i][j])
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

func MIN(a, b int) int {
	if a < b {
		return a
	}
	return b
}
