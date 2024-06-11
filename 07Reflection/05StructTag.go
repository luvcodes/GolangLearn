package main

import "reflect"

type resume struct {
	Name string `info:"name" doc:"我的名字`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		println("info: ", tagInfo, "doc: ", tagDoc)
	}
}

func main() {
	// 使用反射来解决结构体标签中的属性
	var re resume

	findTag(&re)
}
