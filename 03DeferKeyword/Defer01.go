package main

import "fmt"

// func 1
func func1() {
	fmt.Println("func1...")
}

// func 2
func func2() {
	fmt.Println("func2...")
}

// func 3
func func3() {
	fmt.Println("func3...")
}

// defer是类似于压栈的形式, 先写的先入栈
func main() {
	defer func1()
	defer func2()
	defer func3()
}
