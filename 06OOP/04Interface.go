package main

// 定义接口 实现多态
type AnimalInterface interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

// 实现接口具体的Cat类
type Cat struct {
	color string
}

// 实现接口具体的Dog类
type Dog struct {
	color string
}

// 实现接口全部方法
func (this *Cat) Sleep() {
	println("Cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

func (this *Dog) Sleep() {
	println("Dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// 接口作为参数，实现多态
func showAnimal(animal AnimalInterface) {
	animal.Sleep()
	println("Color:", animal.GetColor())
	println("Type:", animal.GetType())
}

func main() {
	// var animal AnimalInterface // 定义接口变量，父类指针
	// animal = &Cat{"White"}
	// animal.Sleep()

	// animal = &Dog{"Black"}
	// animal.Sleep()

	cat := Cat{"White"}
	dog := Dog{"Black"}

	showAnimal(&cat)
	showAnimal(&dog)

}
