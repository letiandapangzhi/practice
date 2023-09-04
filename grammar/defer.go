package main

import (
	"fmt"
)

func test() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

func test1() (i int) {
	i = 0
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

func main() {
	a := test()
	fmt.Println("返回结果", a) // 0

	a1 := test1()
	fmt.Println("返回结果", a1) // 1
}
