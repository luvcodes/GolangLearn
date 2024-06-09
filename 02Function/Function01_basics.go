package main

import "fmt"

// 定义函数
func foo1(a string, b int) int {
	fmt.Println(a)
	fmt.Println(b)

	c := 100
	return c
}

// 两个返回值 (匿名)
func foo2(c string, d int) (int, int) {
	fmt.Println(c)
	fmt.Println(d)

	return 200, 300
}

// 多个返回值 (有形参名)
func foo3(e string, f int) (r1 int, r2 int) {
	fmt.Println("------foo3------")
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)

	// r1 r2 属于foo3的形参 默认值为0
	// r1 r2的作用域空间是整个foo3 {} 函数体空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	r1 = 1000
	r2 = 2000
	return r1, r2
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("------foo4------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 1000
	r2 = 2000
	return r1, r2
}

func main() {
	// 调用函数
	c := foo1("hello", 100)
	fmt.Println(c)

	// 返回两个形参
	result1, result2 := foo2("hello", 100)
	fmt.Println("result1:", result1, "result2:", result2)

	// 对应foo3 给返回值定义名称
	ret1, ret2 := foo3("hello", 100)
	fmt.Println("ret1:", ret1, "ret2:", ret2)

	// 多形参 相同返回类型
	ret3, ret4 := foo4("hello", 100)
	fmt.Println("ret3:", ret3, "ret4:", ret4)
}
