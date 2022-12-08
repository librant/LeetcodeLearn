package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目描述：多线程实现循环打印 123

func main() {

	// 定义三个信号
	signal1 := make(chan struct{})
	signal2 := make(chan struct{})
	signal3 := make(chan struct{})

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("%d", 1)
			// 向 channel 1 中写入数据
			signal1 <- struct{}{}
			// 等待 channel 3 中的数据
			<- signal3
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			time.Sleep(500 * time.Millisecond)
			// 读取 channel 1 中的数据
			<- signal1
			fmt.Printf("%d", 2)
			// 向 channel 2 中写入数据
			signal2 <- struct{}{}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			time.Sleep(500 * time.Millisecond)
			// 读取 channel 2 中的数据
			<- signal2
			fmt.Printf("%d", 3)
			signal3 <- struct{}{}
		}
	}()

	wg.Wait()
}




