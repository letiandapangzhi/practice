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

func deferTest3(i, v int) int {
	fmt.Println(i)
	return i + 1
}

type TestType struct{}

func (t *TestType) deferTest(i int) *TestType {
	fmt.Println(i)
	return t
}

func main() {
	a := deferTest()
	fmt.Println("返回结果", a) // 0

	a1 := deferTest1()
	fmt.Println("返回结果", a1) // 1

	defer deferTest3(1, deferTest3(2, 0)) // 2 1
	// 压栈一次，为了得到第二个参数，需要先计算出deferTest3(2, 0)返回值

	a2 := TestType{}
	defer a2.deferTest(3).deferTest(4) // 3 4
	// 压栈一次，为了得到第二个函数，需要先计算出a2.deferTest(3)返回值
}
