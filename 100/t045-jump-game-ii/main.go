package main

import "fmt"

// https://leetcode.cn/problems/jump-game-ii/

func main() {
	nums := []int{2,3,1,1,4}

	fmt.Printf("%v\n", jump(nums))
}

// 记录跳跃最少的次数，则需要每次跳跃的距离最大
// 贪心算法，每次找到最大的跳跃位置
func jump(nums []int) int {
	n := len(nums)
	maxPos := 0
	end := 0
	res := 0
	// 到达最后时，不需要再跳跃
	for i := 0; i < n - 1; i++{
		// 获取当前位置能跳跃的最大距离，
		maxPos = MAX(maxPos, nums[i]+i)
		// 每次找到能跳跃到的最远边界，最远边界之后，再遍历其中能到达的最远边界
		if i == end {
			end = maxPos
			// 到达边界时，跳跃增加 1
			res++
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



