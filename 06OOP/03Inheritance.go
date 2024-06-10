package main

import "fmt"

// 实现继承
type Human struct {
	Name   string
	Gender string
}

type SuperHuman struct {
	Human
	Level int
}

func (this *Human) Eat() {
	fmt.Println("Human is eating")
}

func (this *Human) Walk() {
	fmt.Println("Human is walking")
}

// =====================================

// 子类重写父类方法
func (this *SuperHuman) Eat() {
	fmt.Println("SuperHuman is eating")
}

func (this *SuperHuman) Walk() {
	fmt.Println("SuperHuman is walking")
}

// 子类新增方法
func (this *SuperHuman) Fly() {
	fmt.Println("SuperHuman is flying")
}

func main() {
	human := Human{"Mark", "Male"}
	human.Eat()
	human.Walk()
	fmt.Println("======================")

	// 子类对象定义
	superman := SuperHuman{Human{"lisi", "Male"}, 88}
	superman.Eat()
	superman.Walk()
	superman.Fly()

	// 子类对象定义第二种形式
	var superman2 SuperHuman
	superman2.Name = "wangwu"
	superman2.Gender = "Male"
	superman2.Level = 99

	superman2.Eat()
	superman2.Walk()
	superman2.Fly()
}
