package main

import "fmt"

// https://leetcode.cn/problems/determine-color-of-a-chessboard-square/

func main() {
	coordinates1 := "a1"
	fmt.Printf("%v\n", squareIsWhite(coordinates1))

	coordinates2 := "h3"
	fmt.Printf("%v\n", squareIsWhite(coordinates2))
}

// 根据规律：直接按照奇偶判断即可
func squareIsWhite(coordinates string) bool {
	n := len(coordinates)
	if n != 2 {
		return false
	}

	if coordinates[0] % 2 == 0 && coordinates[1] % 2 == 0 {
		return false
	}
	if coordinates[0] % 2 == 1 && coordinates[1] % 2 == 1 {
		return false
	}
	if coordinates[0] % 2 == 1 && coordinates[1] % 2 == 0 {
		return true
	}
	if coordinates[0] % 2 == 0 && coordinates[1] % 2 == 1 {
		return true
	}
	return false
}
