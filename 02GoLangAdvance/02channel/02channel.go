package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个缓冲区为3的channel
	c := make(chan int, 3)

	// 开启一个goroutine
	go func() {
		defer fmt.Println("go routine finished")

		// 向channel中写入数据
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("go routine write ", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()

	// 休眠两秒
	time.Sleep(2 * time.Second)

	// 从channel中读取数据
	for i := 0; i < 3; i++ {
		// 定义变量接收channel中的数据
		num := <-c
		fmt.Println("num = ", num)

	}

	fmt.Println("main finished")
}
