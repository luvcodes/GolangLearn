package main

// import "fmt"

// func printArray(myArray [4]int) {
// 	// 值拷贝
// 	for index, value := range myArray {
// 		fmt.Printf("index: %d, value: %d\n", index, value)
// 	}

// 	// 因为这种数组直接定义长度的方式,就是值拷贝,即使修改了某个元素也只是在这个函数体内有效
// 	myArray[0] = 111
// }

// func main() {
// 	var myArray1 [10]int
// 	myArray2 := [10]int{1, 2, 3, 4}
// 	myArray3 := [4]int{11, 22, 33, 44}

// 	//// 遍历数组
// 	//for i := 0; i < len(myArray1); i++ {
// 	//	fmt.Println(myArray1[i])
// 	//}
// 	//
// 	//// 使用range关键字给出数组元素下标和元素值
// 	//for index, value := range myArray2 {
// 	//	fmt.Println("index = ", index, "value = ", value)
// 	//}

// 	// 查看三个array的类型
// 	fmt.Printf("myArray1 types = %T\n", myArray1)
// 	fmt.Printf("myArray2 types = %T\n", myArray2)
// 	fmt.Printf("myArray3 types = %T\n", myArray3)

// 	// 数组传参必须要长度匹配
// 	// printArray(myArray1) // 报错: Cannot use 'myArray1' (type [10]int) as the type [4]int
// 	printArray(myArray3)

// }
