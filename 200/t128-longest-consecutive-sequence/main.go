package main

import "fmt"

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

func main() {
	nums := []int{100,4,200,1,3,2}

	fmt.Printf("%d\n", longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	n := len(nums)
	arrMap := make(map[int]bool)
	for i := 0; i < n; i++ {
		// 将所有的数据作为 key 存入
		arrMap[nums[i]] = true
	}
	maxCnt := 0
	for k := range arrMap {
		// 左边的值存在则，则直接进入下一个循环，因为在下一次遍历的时候，需要从左值开始遍历
		if !arrMap[k-1] {
			// 记录当前的 k 的值，这个就是连续的起始值
			currentNum := k
			// 当前连续的第一个值
			currentStreak := 1
			// 再往前进行遍历，看连续存在的长度
			for arrMap[currentNum+1] {
				currentNum++
				currentStreak++
			}
			// 返回当前最大的连续值
			maxCnt = MAX(maxCnt, currentStreak)
		}
	}
	return maxCnt
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}
