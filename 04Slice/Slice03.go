package main

import "fmt"

func main() {
	// 声明slice1是一个切片,初始化,默认值为1,2,3, 长度len是3
	slice1 := []int{1, 2, 3}

	// 声明slice2是一个切片,但是并没有给slice2分配空间
	var slice2 []int
	// 给slice2开辟3个空间,默认值是0
	//slice2 = make([]int, 3)

	fmt.Printf("len = %d, slice1: %v\n", len(slice1), slice1)
	fmt.Printf("len = %d, slice2: %v", len(slice2), slice2)
}
