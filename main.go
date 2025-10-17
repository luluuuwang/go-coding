package main

// 启动10个goroutine，他们各自打印自己的编号，最终输出顺序必须是 1 - 10

// 通过channel控制下一个goroutine何时可以打印

// 10.17

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	chs := make([]chan struct{}, 12)

	// 创建十个信号通道
	for i := 0; i < 10; i++ {
		chs[i] = make(chan struct{})
	}

	// 创建十个goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) { // 定义参数
			defer wg.Done()
			<-chs[i] // 没接收到消息时，就一直阻塞
			fmt.Println(i + 1)

			if i+1 < 10 {
				chs[i+1] <- struct{}{}
			}
		}(i) // 传入当前i
	}

	chs[0] <- struct{}{}
	wg.Wait()
}
