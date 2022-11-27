package main

import "fmt"

// https://leetcode.cn/problems/number-of-islands/

// bfs: 四个方向都需要进行遍历

func main() {
	grid := [][]byte{
	{'1', '1', '1', '1', '0'},
	{'1', '1', '0', '1', '0'},
	{'1', '1', '0', '0', '0'},
	{'0', '0', '0', '0', '0'},
	}

	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, r, c int) {
	h, w := len(grid), len(grid[0])
	if r < 0 || r >= h || c < 0 || c >= w {
		return
	}
	// 当前节点不是 1 则表示不是大陆
	if grid[r][c] != '1' {
		return
	}
	// 避免重复搜索
	grid[r][c] = '2'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}
