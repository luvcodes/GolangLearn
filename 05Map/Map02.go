package main

import "fmt"

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "New York"

	printMap(cityMap)

	// 删除
	delete(cityMap, "China")

	// 修改
	cityMap["USA"] = "DC"
	changeMap(cityMap)
	fmt.Println("=========================")

	// 再次遍历
	printMap(cityMap)
	fmt.Println("=========================")

	// 调用拷贝map的方法
	printMap(copyMap(cityMap))
}

// 参数传递map是引用传递，传递的是一个指针
func printMap(cityMap map[string]string) {
	// 遍历cityMap
	for key, value := range cityMap {
		fmt.Println("key:", key, "value:", value)
	}
}

func changeMap(cityMap map[string]string) {
	cityMap["UK"] = "London"
}

// 完全拷贝map
func copyMap(cityMap map[string]string) map[string]string {
	newMap := make(map[string]string)
	// 遍历cityMap依次赋值
	for key, value := range cityMap {
		newMap[key] = value
	}
	fmt.Println("Copy map below: ")
	return newMap

}
