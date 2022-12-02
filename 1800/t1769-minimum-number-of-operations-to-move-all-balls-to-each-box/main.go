package main

import "fmt"

func main() {
	boxes1 := "110"
	fmt.Printf("%v\n", minOperations(boxes1))

	boxes2 := "001011"
	fmt.Printf("%v\n", minOperations(boxes2))
}

// 暴力遍历
func minOperations(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		// 先计算从 0 ~ i-1 位置的
		count := 0
		for j := 0; j < i; j++ {
			if boxes[j] == '1' {
				count += (i - j)
			}
		}
		// 再计算从 i+1 ~ n-1 位置的数
		for k := i + 1; k < n; k++ {
			if boxes[k] == '1' {
				count += (k - i)
			}
		}
		res[i] = count
	}
	return res
}
