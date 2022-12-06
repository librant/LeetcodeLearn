package main

import (
	"fmt"
	"unicode"
)

// https://leetcode.cn/problems/number-of-different-integers-in-a-string/

func main() {
	word1 := "a123bc34d8ef34"
	fmt.Printf("%v\n", numDifferentIntegers(word1))

	word2 := "leet1234code234"
	fmt.Printf("%v\n", numDifferentIntegers(word2))

	word3 := "a1b01c001"
	fmt.Printf("%v\n", numDifferentIntegers(word3))
}

// 双指针，先确定左边指针，再确定右边指针，需要去掉前导 0

func numDifferentIntegers(word string) int {
	n := len(word)
	s := make(map[string]bool)
	p1 := 0
	for {
		for p1 < n && !unicode.IsDigit(rune(word[p1])) {
			p1++
		}
		if p1 == n {
			break
		}
		p2 := p1
		for p2 < n && unicode.IsDigit(rune(word[p2])) {
			p2++
		}
		for p2-p1 > 1 && word[p1] == '0' { // 去除前导 0
			p1++
		}
		s[word[p1:p2]] = true
		p1 = p2
	}
	return len(s)
}


