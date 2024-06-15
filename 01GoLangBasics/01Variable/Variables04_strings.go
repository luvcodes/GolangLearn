package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 拼接字符串的3种方式

	// 1. 使用 + 进行拼接
	str1 := "hello " +
		"world"
	fmt.Println(str1)
	fmt.Println("=====================================")

	// 2. 使用buffer进行拼接
	str2 := "hello "
	str3 := "world"
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(str2)
	stringBuilder.WriteString(str3)
	fmt.Println(stringBuilder.String())
	fmt.Println("=====================================")

	// 3. 使用fmt.Sprintf进行拼接
	str4 := fmt.Sprintf("%s%s", str2, str3)
	fmt.Println(str4)
}
