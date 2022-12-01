package main

import "fmt"

// https://leetcode.cn/problems/maximum-frequency-stack/

// 核心思路：把频率（出现次数）不同的元素，压入不同的栈中。每次出栈时，弹出含有频率最高元素的栈的栈顶。

func main() {
	fs := Constructor()

	fs.Push (5)
	fmt.Printf("%v\n", fs)
	fs.Push (7)
	fmt.Printf("%v\n", fs)
}

type FreqStack struct {
	cnt    map[int]int
	stacks [][]int
}

func Constructor() FreqStack {
	return FreqStack{cnt: make(map[int]int)}
}

func (this *FreqStack) Push(val int) {
	c := this.cnt[val]
	if c == len(this.stacks) { // 这个元素的频率已经是目前最多的，现在又出现了一次
		this.stacks = append(this.stacks, []int{val}) // 那么必须创建一个新栈
	} else {
		this.stacks[c] = append(this.stacks[c], val) // 否则就压入对应的栈
	}
	this.cnt[val]++ // 更新频率
}

func (this *FreqStack) Pop() int {
	back := len(this.stacks) - 1
	st := this.stacks[back]
	bk := len(st) - 1
	val := st[bk] // 弹出最右侧栈的栈顶
	if bk == 0 { // 栈为空
		this.stacks = this.stacks[:back] // 删除
	} else {
		this.stacks[back] = st[:bk]
	}
	this.cnt[val]-- // 更新频率
	return val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
