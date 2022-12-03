package main

import (
	"fmt"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("%v\n", maxProfitTimeOut(prices1))

	prices2 := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("%v\n", maxProfit(prices2))
}

// o(n) 的时间复杂度
func maxProfit(prices []int) int {
	// 在遍历的时候，记录一个历史的最低价格，当前卖出的时候，只需要知道历史最低价格就可以
	n := len(prices)
	if n <= 1 {
		return 0
	}
	minPrice := prices[0]
	res := 0
	for i := 1; i < n; i++ {
		minPrice = MIN(minPrice, prices[i])
		res = MAX(prices[i]-minPrice, res)
	}
	return res
}

// 第一种 o(n2) 直接遍历
func maxProfitTimeOut(prices []int) int {
	n := len(prices)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res = MAX(res, prices[j]-prices[i])
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

func MIN(a, b int) int {
	if a < b {
		return a
	}
	return b
}
