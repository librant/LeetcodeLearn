package main

import "fmt"

// https://leetcode.cn/problems/permutations/

func main() {
	nums := []int{1,2,3}

	fmt.Printf("%v\n", permute(nums))
}

// 回溯算法：
// nums: 原始列表
// pathNums: 路径上的数字
// used: 是否访问过

func permute(nums []int) (res [][]int) {
	n := len(nums)
	used := make([]bool, n)
	var pathNums []int
	var backtrace func(nums, pathNums []int, used []bool)
	backtrace = func (nums, pathNums []int, used []bool) {
		// 这里的结束条件, 当数组中所有的数字都被选择了之后
		if n == len(pathNums) {
			// 此时 pathNums 就是已经排好的
			// 这个地方需要用新的切片拷贝才行，可能容器疏忽
			tmp := make([]int, len(nums))
			// 这里是因为 pathNums 是共用的，下一次的循环可以改变 pathNums 的值，需要用新的切片拷贝
			copy(tmp, pathNums)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			// 当前这个位置的数还没有被遍历到，则可以加入到 pathNums，继续开始回溯下一次
			if !used[i] {
				pathNums = append(pathNums, nums[i])
				// 继续往下查找下一个没有被选择的数据
				used[i] = true
				backtrace(nums, pathNums, used)
				// 这里这次回溯完成，需要等待下一次的回溯
				pathNums = pathNums[:len(pathNums) -1]
				used[i] = false
			}
		}
	}
	backtrace(nums, pathNums, used)
	return
}
