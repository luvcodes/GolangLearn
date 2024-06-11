package main

import "fmt"

func main() {
	// 1. Map的第一种声明方式
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is nil")
	}

	// 使用make为myMap1分配内存空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "python"
	myMap1["three"] = "ruby"
	myMap1["four"] = "go"

	// 插入顺序不是存储顺序
	fmt.Println(myMap1)

	// 2. Map的第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "python"
	myMap2[3] = "ruby"
	myMap2[4] = "go"
	fmt.Println(myMap2)

	// 3. Map的第三种声明方式
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "C++",
		"three": "Go",
		"four":  "JavaScript",
	}
	fmt.Println(myMap3)
}
