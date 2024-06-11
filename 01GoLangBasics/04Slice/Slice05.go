package main

import "fmt"

func main() {
	// 切片截取
	s := []int{1, 2, 3} // len = 3, cap = 3, [1,2,3]
	s1 := s[0:2]        // [1,2] // 从0开始，截取2个元素
	for index, value := range s {
		fmt.Printf("s[%d] = %d\n", index, value)
	}
	fmt.Println("===============")

	s1[0] = 100 // 修改s1的第1个元素，两个数组的值都会发生改变，因为指向的是同一个数组
	fmt.Println("s = ", s)
	fmt.Println("s1 = ", s1)

	// 使用copy方法 将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) // 开辟3个容量 [0,0,0]
	copy(s2, s)          // 将s中的值依次拷贝到s2中
	fmt.Println("s2 = ", s2)

}
