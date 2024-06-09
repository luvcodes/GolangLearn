package main

import "fmt"

// 使用引用传递
func printArray2(myArray []int) {
	// _ 是占位符
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	// 引用传递的方式,修改数组元素后,就会彻底修改元素值
	myArray[0] = 111
}

func main() {
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("myArray = %v\n", myArray)

	printArray2(myArray)

	fmt.Println("==========")

	// 再次遍历数组，查看修改过后的数组元素值
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
