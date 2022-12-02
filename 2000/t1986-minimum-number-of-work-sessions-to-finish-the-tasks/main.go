package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-work-sessions-to-finish-the-tasks/

func main() {
	tasks := []int{3,1,3,1,1}
	sessionTime := 8
	fmt.Printf("%v\n", minSessions(tasks, sessionTime))
}

func minSessions(tasks []int, sessionTime int) int {
	// 先对数组进行排序
	sort.Ints(tasks)




	return 0
}
