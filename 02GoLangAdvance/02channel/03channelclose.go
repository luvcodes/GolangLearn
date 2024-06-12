package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close可以关闭一个channel，如果不关闭后面的主协程就会阻塞
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main goroutine ends...")
}
