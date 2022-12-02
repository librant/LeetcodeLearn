package main

import "fmt"

// https://leetcode.cn/problems/combination-sum/

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7

	fmt.Printf("%v\n", combinationSum(candidates, target))
}

// 递归实现：
// 需要从当前的位置进行遍历，每次递归需要加上当前值作为 sum + 当前值作为下一次的 递归调用

func combinationSum(candidates []int, target int) (ans [][]int) {
	var path []int
	var traverse func(sum, startIndex int)
	traverse = func(sum, startIndex int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			if sum+candidates[i] > target { // 剪枝
				continue
			}
			path = append(path, candidates[i])
			// 需要包含当前的 i 的位置
			traverse(sum+candidates[i], i) // 不走回头路
			path = path[:len(path)-1]
		}
	}
	traverse(0, 0)
	return
}
