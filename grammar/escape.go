package main

import (
	"strings"
)

func main() {
	// go build -gcflags="-m" escape.go
	// 局部变量通常分配在栈上，当编译器分析到外部存在引用会内存逃逸分配到堆上
	// 返回引用类型会出现内存逃逸，返回值类型不会，个人理解是值类型是一个拷贝
	//string1 := escapeTest1()
	//strings.Split(string1, ",")
	//
	//escapeTest2()
	//
	//map1 := escapeTest3()
	//if _, exist := map1["test"]; exist {
	//
	//}

	slice1 := make([]string, 100000)
	println(slice1)
}

func escapeTest1() string {
	var slice1 = []string{"1", "2", "3", "4", "5"}
	return strings.Join(slice1, ",")
}

func escapeTest2() []string {
	return []string{"1", "5"}
}
func escapeTest3() map[string]string {
	return make(map[string]string)
}
