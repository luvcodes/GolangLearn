package main

import "fmt"

func deferFunction() int {
	fmt.Println("deferFunction...")
	return 0
}

func returnFunction() int {
	fmt.Println("returnFunction...")
	return 0
}

func returnAndDefer() int {
	defer deferFunction()
	return returnFunction()
}

// 先运行return语句,再执行defer语句
func main() {
	returnAndDefer()
}
