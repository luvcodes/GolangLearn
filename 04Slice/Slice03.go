package main

// import "fmt"

// func main() {
// 	// 声明切片方式1. 声明slice1是一个切片,初始化,默认值为1,2,3
// 	slice1 := []int{1, 2, 3}

// 	// 声明切片方式2. 声明slice2是一个切片,但是并没有给slice2分配空间
// 	var slice2 []int
// 	// 给slice2开辟3个空间,默认值是0
// 	slice2 = make([]int, 3)

// 	// 声明切片方式3. 声明slice3是一个切片，同时给slice3分配空间，3个空间，初始化为0
// 	var slice3 []int = make([]int, 3)

// 	// 声明切片方式4. 针对3的改进，通过:=推导出slice4是一个切片
// 	slice4 := make([]int, 3)

// 	fmt.Printf("len = %d, slice1: %v\n", len(slice1), slice1)
// 	fmt.Printf("len = %d, slice2: %v", len(slice2), slice2)
// 	fmt.Printf("len = %d, slice3: %v\n", len(slice3), slice3)
// 	fmt.Printf("len = %d, slice4: %v\n", len(slice4), slice4)

// 	// 判断一个slice是否为空
// 	if slice2 == nil {
// 		fmt.Println("slice2 is nil")
// 	} else {
// 		fmt.Println("slice2 is not nil")
// 	}
// }
