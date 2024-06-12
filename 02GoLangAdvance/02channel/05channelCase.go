package main

import "fmt"

// identify fibonacci function
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// 发送数据到channel c
		case c <- x:
			x, y = y, x+y
		// 从channel quit中取数据, 如果quit channel中有数据 (0)，说明子协程中的任务已经完成，退出
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// 定义两个channel
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c) // 从channel c中取数据
		}

		quit <- 0 // 发送数据到channel quit
	}()

	fibonacci(c, quit)
}
