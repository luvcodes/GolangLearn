package main

import "fmt"

// 定义一种新的数据类型myint，是int的一个别名
type myint int

// Book 定义一个结构体
type Book struct {
	title string
	auth  string
}

func main() {
	// 使用自定义数据类型
	//var a myint = 10
	//fmt.Println("a = ", a)
	//fmt.Printf("type of %T\n", a)

	var book1 Book
	book1.title = "Golang intro"
	book1.auth = "123456"
	fmt.Println("book1:", book1)

	//changeBook(book1)
	//fmt.Println(book1)

	actualChangeBook(&book1)
	fmt.Println("book2:", book1)

}

// 这样写就是传递一个book的副本，并没有真正改变book1的值
func changeBook(book Book) {
	book.title = "Golang advance"
}

func actualChangeBook(book *Book) {
	book.title = "Golang advance"
}
