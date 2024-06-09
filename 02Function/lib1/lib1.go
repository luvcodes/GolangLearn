package lib1

import "fmt"

// 首字母大写代表是public方法
func Lib1Test() {
	fmt.Println("lib1Test API...")
}

func init() {
	fmt.Println("lib1 init")
}
