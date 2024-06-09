package main

// import "fmt"

// func main() {
// 	// 切片容量的追加
// 	// len = 3 cap = 5
// 	var numbers = make([]int, 3, 5)

// 	// 追加一个元素
// 	numbers = append(numbers, 1)
// 	fmt.Println(numbers, len(numbers), cap(numbers)) // [0 0 0 1] 4 5

// 	// 追加一个元素
// 	numbers = append(numbers, 2)
// 	fmt.Println(numbers, len(numbers), cap(numbers)) // [0 0 0 1 2] 5 5

// 	// 达到了cap的容量，再追加一个元素，会重新分配内存，容量变为原来的2倍
// 	numbers = append(numbers, 3)
// 	fmt.Println(numbers, len(numbers), cap(numbers)) // [0 0 0 1 2 3] 6 10

// 	fmt.Println("=====================================")
// 	var numbers2 = make([]int, 3)                       // 没有指定cap，默认和len一样
// 	fmt.Println(numbers2, len(numbers2), cap(numbers2)) // [0 0 0] 3 3
// 	numbers2 = append(numbers2, 1)
// 	fmt.Println(numbers2, len(numbers2), cap(numbers2)) // [0 0 0 1] 4 6 // cap = 6 = 3 * 2 自动扩容为原来cap的2倍
// }
