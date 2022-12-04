package main

import "fmt"

// https://leetcode.cn/problems/closest-dessert-cost/

func main() {
	baseCosts := []int{1,7}
	toppingCosts := []int{3,4}
	target := 10

	fmt.Printf("%v\n", closestCost(baseCosts, toppingCosts, target))
}

// 动态规划：
// dp[i]: 表示制作甜品开销为 i 是否存在合法方案
// 对于 baseCosts 中的开销均为合法 dp[baseCosts[i]] = true  baseCosts[i] < target
// 每种配料最多增加两种：需要放置的配置开销为 j
// dp[i] = dp[i-j] | dp[i] (i > j)
// 这里的动态规划比较难想到

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	n := len(baseCosts)
	if n == 0 {
		return 0
	}
	// 当前基料中的最小值
	x := baseCosts[0]
	for i := 1; i < n; i++ {
		x = MIN(x, baseCosts[i])
	}
	// 如果最小的基料都超过 target, 则返回基料最小值即可
	if x >= target {
		return x
	}
	dp := make([]bool, target+1)
	ans := 2*target - x
	for i := 0; i < n; i++ {
		if baseCosts[i] <= target {
			dp[baseCosts[i]] = true
		} else {
			ans = MIN(ans, baseCosts[i])
		}
	}

	for _, c := range toppingCosts {
		for count := 0; count < 2; count++ {
			for t := target; t > 0; t-- {
				if dp[t] && t + c > target {
					ans = MIN(ans, t+c)
				}
				if t - c > 0 {
					dp[t] = dp[t] || dp[t-c]
				}
			}
		}
	}

	for i := 0; i <= ans-target; i++ {
		if dp[target-i] {
			return target-i
		}
	}
	return ans
}

func MIN(a, b int) int {
	if a < b {
		return a
	}
	return b
}


