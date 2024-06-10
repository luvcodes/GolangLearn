package main

import "fmt"

// interface{} 表示空接口，可以接收任何类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// interface{}如何区分具体类型?
	// 使用类型断言
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type")
		fmt.Printf("value is %v\n", value)
	}
}

type book struct {
	name string
}

func main() {
	book := book{"Golang"}
	myFunc(book)

	str1 := "aaa"
	myFunc(str1)

	integer1 := 100
	myFunc(integer1)

	float1 := 3.14
	myFunc(float1)
}
