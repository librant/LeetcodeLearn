package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/

func main() {
	x := 3
	y := 4
	points := [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}

	fmt.Printf("%v\n", nearestValidPoint(x, y, points))
}

func nearestValidPoint(x int, y int, points [][]int) int {
	res := math.MaxInt32
	min := -1
	for i := 0; i < len(points); i++ {
		// 如果有 x 或者 y 相同，则为符合条件的点
		if points[i][0] == x || points[i][1] == y {
			// 当前的曼哈顿距离最小, 则记录当前的最小值
			dist := ABS(x-points[i][0]) + ABS(y-points[i][1])
			if dist < res {
				res = dist
				// 这里返回最小 i 的坐标值即可
				min = i
			}
		}
	}
	return min
}

func ABS(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
