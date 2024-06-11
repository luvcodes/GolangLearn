package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("go routine finished")

		fmt.Println("go routine started")

		c <- 666
	}()

	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("main finished")

}
