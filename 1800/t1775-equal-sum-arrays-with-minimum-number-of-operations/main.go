package main

import "fmt"

// https://leetcode.cn/problems/equal-sum-arrays-with-minimum-number-of-operations/

func main() {
	nums1 := []int{1,2,3,4,5,6}
	nums2 := []int{1,1,2,2,2,2}

	fmt.Printf("%v\n", minOperations(nums1, nums2))
}

func help(h1 [7]int, h2 [7]int, diff int) (res int) {
	h := [7]int{}
	for i := 1; i < 7; i++ {
		h[6-i] += h1[i]
		h[i-1] += h2[i]
	}
	for i := 5; i > 0 && diff > 0; i-- {
		t := min((diff+i-1)/i, h[i])
		res += t
		diff -= t * i
	}
	return res
}

func minOperations(nums1 []int, nums2 []int) (ans int) {
	n, m := len(nums1), len(nums2)
	if 6*n < m || 6*m < n {
		return -1
	}
	var cnt1, cnt2 [7]int
	diff := 0
	for _, i := range nums1 {
		cnt1[i]++
		diff += i
	}
	for _, i := range nums2 {
		cnt2[i]++
		diff -= i
	}
	if diff == 0 {
		return 0
	}
	if diff > 0 {
		return help(cnt2, cnt1, diff)
	}
	return help(cnt1, cnt2, -diff)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}