package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("交换前，a 的值 : %d, b的值 : %d\n", a, b)
	// swap(a, b) // 这种只能是值传递
	swap(&a, &b)
	fmt.Printf("交换后，a 的值 : %d, b的值 : %d\n", a, b)
}

// 这样是值传递，不会改变原来的值
// 值传递的概念主要就是在于，传递的是值的副本，而不是值本身
// 也就是说，传递的是值的副本，对副本的修改，也就是在函数内修改参数的值，不会影响原来的值
//func swap(x, y int) int {
//	var temp int
//
//	temp = x /* 保存 x 的值 */
//	x = y    /* 将 y 值赋给 x */
//	y = temp /* 将 temp 值赋给 y */
//
//	return temp
//}

// 这样是引用传递，会改变原来的值
// 但是，如果传递的是指针，那么就是引用传递，对指针的修改，也就是在函数内修改参数的值，会影响原来的值
func swap(pointerA *int, pointerB *int) {
	var temp int

	temp = *pointerA
	*pointerA = *pointerB
	*pointerB = temp

}
