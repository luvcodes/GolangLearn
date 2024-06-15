package main

import "fmt"

func main() {
	a := 200
	b := 100

	// swap the value of these two variables using the third variable
	// var temp int
	// temp = a
	// a = b
	// b = temp

	// swap the value of these two variables using a second function (pointer)
	swap(&a, &b)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}

func swap(a *int, b *int) {
	// implement the swap function using pointers
	// identify the third variable temp
	temp := *a
	*a = *b
	*b = temp
}
