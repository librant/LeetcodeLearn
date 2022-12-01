package main

import "fmt"

// https://leetcode.cn/problems/jump-game/?favorite=2cktkvj

// 这题是要求能到达，不一定是刚好到达，只要最后一个坐标大于 n-1，则可以到达

func main() {
	nums := []int{2,3,1,1,4}

	fmt.Printf("%v\n", canJumpTimeout(nums))
	fmt.Printf("%v\n", canJump(nums))
}

func canJump(nums []int) bool {
	n := len(nums)
	pos := 0
	for i := 0; i < n; {
		// 每次能到达的最大位置，只要大于 n-1 的位置，则认为是可以到达最后
		pos = MAX(pos, i+nums[i])
		if pos >= n-1 {
			return true
		}
		i = pos
	}
	return false
}

func canJumpTimeout(nums []int) bool {
	n := len(nums)
	pos := 0
	for i := 0; i < n; i++{
		// 直接从开始位置往后跳, 当前索引的位置最远的位置大，则不可能跳到最后，顺着跳超时
		if i > pos {
			return false
		}
		pos = MAX(pos, i+nums[i])
	}
	return true
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}


