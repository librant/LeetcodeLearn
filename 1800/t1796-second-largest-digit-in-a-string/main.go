package main

import (
	"fmt"
	"unicode"
)

// https://leetcode.cn/problems/second-largest-digit-in-a-string/

func main() {
	s1 := "dfa12321afd"
	fmt.Printf("%v\n", secondHighest(s1))
	s2 := "abc1111"
	fmt.Printf("%v\n", secondHighest(s2))
	s3 := "ck077"
	fmt.Printf("%v\n", secondHighestNew(s3))
}

// 返回字符串中第二大的数字

func secondHighestNew(s string) int {
	first, second := -1, -1
	for _, c := range s {
		if unicode.IsDigit(c) {
			num := int(c - '0')
			if num > first {
				second = first
				first = num
			} else if second < num && num < first {
				second = num
			}
		}
	}
	return second
}

func secondHighest(s string) int {
	// 维护一个数组，这个数组总是排前两个最大的数
	var ss []int
	for i := 0; i < len(s); i++ {
		// 查找其中的数字
		if s[i] >= '0' && s[i] <= '9' {
			// 当前 ss 的长度超过 2 个，则需要维护一个不相等的数组
			if len(ss) >= 2 {
				// 如果有数相等， 则忽略
				if int(s[i]-'0') == ss[0] || int(s[i]-'0') == ss[1] {
					continue
				}
				// 如果当前的数大于第一个数，则需要交换位置
				if int(s[i]-'0') > ss[0] {
					temp := ss[0]
					ss[0] = int(s[i]-'0')
					ss[1] = temp
					continue
				}
				// 如果当前的数大于第二个数，则需要交换位置
				if int(s[i]-'0') > ss[1] {
					ss[1] = int(s[i]-'0')
					continue
				}
			}
			if len(ss) == 1 {
				if int(s[i]-'0') == ss[0] {
					// 如果相等就不加入到 ss
					continue
				}
				ss = append(ss, int(s[i]-'0'))
				if ss[0] < ss[1] {
					temp := ss[0]
					ss[0] = ss[1]
					ss[1] = temp
				}
			}
			if len(ss) == 0 {
				ss = append(ss, int(s[i]-'0'))
			}
		}
	}

	if len(ss) < 2 {
		return -1
	}
	return ss[1]
}
