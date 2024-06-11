package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book")
}

func main() {
	// b: pair<type: Book, value: &Book{}地址>
	book := &Book{}

	// r: pair<type:, value:>
	var r Reader
	// r: pair<type: *Book, value: &Book{}地址>
	r = book
	r.ReadBook()

	// w: pair<type:, value:>
	var w Writer
	// r: pair<type: *Book, value: &Book{}地址>
	w = r.(Writer) // 此处断言为什么会成功？因为 w r 具体的type是一致的
	w.WriteBook()

}
