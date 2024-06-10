package main

import "fmt"

// Hero 定义Hero类，定义setter和getter方法
type Hero struct {
	Name  string
	Id    int
	Level int
}

/*// GetName getter
func (this Hero) GetName() string {
	return this.Name
}

// SetName setter
// this是调用该方法的对象的一个拷贝
func (this Hero) SetName(newName string) {
	this.Name = newName
}*/

func (this *Hero) showHero() {
	fmt.Println("Name is", this.Name)
	// fmt.Println("Id is", this.Id)
	// fmt.Println("Level is", this.Level)
}

// GetName getter
func (this *Hero) GetName() string {
	return this.Name
}

// SetName setter
// this是调用该方法的对象的一个拷贝
func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	// 创建Hero对象
	hero := Hero{Name: "zhangsan", Id: 1, Level: 1}
	hero.showHero()

	hero.SetName("lisi") // 调用set方法
	hero.showHero()

}
