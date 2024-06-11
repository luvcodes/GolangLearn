package main

import "fmt"

// const来定义枚举类型
const (
	aa = 1
	bb = 2
)

// 可以在 const() 添加一个关键字 iota，每行的iota都会累加1，第一行的iota的默认值为0
const (
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

const (
	// iota = 0, a = iota + 1, b = iota + 2
	a, b = iota + 1, iota + 2
	c, d
	e, f
	g, h = iota * 2, iota * 3
	i, k
)

func main() {
	// 常量 (只读属性)
	const length int = 10
	fmt.Println(length)

	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println(SHENZHEN)

	fmt.Println("a = ", a, "b = ", b, "c = ", c, "d = ", d, "e = ", e, "f = ", f)
	fmt.Println("g = ", g, "h = ", h, "i = ", i, "k = ", k)
}
