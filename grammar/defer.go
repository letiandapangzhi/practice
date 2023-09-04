package main

import (
	"fmt"
)

func deferTest() int {
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

func deferTest1() (i int) {
	i = 0
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

func main() {
	a := deferTest()
	fmt.Println("返回结果", a) // 0

	a1 := deferTest1()
	fmt.Println("返回结果", a1) // 1
}
