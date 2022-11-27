package main

import "fmt"

// https://leetcode.cn/problems/sort-colors/?favorite=2cktkvj

func main() {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
	fmt.Printf("%v", nums)
}

func sortColors(nums []int)  {
	// 原地更新，不能用 sort 函数
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			// 如果是 0 则不用进行调换
			continue
		}
		if nums[i] == 1 {
			// 如果是1， 则找最右边第一个为 0 的进行置换， 没有零则不进行置换
			for j := i+1; j < n; j++ {
				if nums[j] == 0 {
					swap(i, j, nums)
					break
				}
			}
		}
		if nums[i] == 2 {
			// 如果是2，则从最右边将 0/1 的数据进行置换, 这里需要优先找 0，没有找到再找 1
			find := false
			for j := i+1; j < n; j++ {
				if nums[j] == 0 {
					swap(i, j, nums)
					find = true
					break
				}
			}
			if find {
				continue
			}
			for j := i+1; j < n; j++ {
				if nums[j] == 1 {
					swap(i, j, nums)
					break
				}
			}
		}
	}
}

func swap(i, j int, nums []int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
