package main

import "fmt"

// https://leetcode.cn/problems/maximal-rectangle/?favorite=2cktkvj

func main() {
	matrix1 := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}

	fmt.Printf("%v\n", maximalRectangleDP(matrix1))

	matrix2 := [][]byte{{'1','0','1','1','1'},{'0','1','0','1','0'},{'1','1','0','1','1'},{'1','1','0','1','1'},{'0','1','1','1','1'}}

	fmt.Printf("%v\n", maximalRectangleDP(matrix2))
}

// 动态规划：（严格意义来说，不算 DP）
// dp[i][j]: 表示第 i 行， 第 j 列 左边 连续 1 的个数，代表以当前坐标矩形的最大宽度
// dp[k][j]: 表示 [k, j] 这个位置最右边的 1 的个数，k 在 [0, i-1] 之间
// 此时，以 [i, j] 为右下角的矩形最大面积则为 dp[k][j] 中与 dp[i][j] 小的值 乘以 矩形的高度 (i-k)

func maximalRectangleDP(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	maxArea := 0
	// 计算每个坐标点的左边连续 1 的个数
	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// 当前位置为 1
			if matrix[i][j] == '1' {
				// 如果是最左边边，则为 1
				if j == 0 {
					dp[i][j] = 1
				} else {
					// 否则为 j-1 列的值 +1
					dp[i][j] = dp[i][j-1] + 1
				}

				// 这里可以直接合并到一个大循环，因为前面的都已经遍历遍历完成，需要计算当前 [i, j] 的最大矩形面积

				// [i, j] 为最右下角的时候，最底层的宽度，上层必须在这个宽度以内
				width := dp[i][j]
				// 第 i 行的面积
				for k := i; k >= 0; k-- {
					// 找到每行的最小值, 只能从 i 行开始遍历， 上面只能用当前 dp[k][j] 的最小值
					width = MIN(width, dp[k][j])
					maxArea = MAX(maxArea, width * (i - k + 1))
				}
			}
		}
	}
	return maxArea
}

func MIN(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}
