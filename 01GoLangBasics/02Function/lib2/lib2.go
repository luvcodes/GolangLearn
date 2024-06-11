package lib2

import "fmt"

// 首字母大写代表是public方法
func Lib2Test() {
	fmt.Println("Lib2Test API...")
}

func init() {
	fmt.Println("lib2 init")
}
