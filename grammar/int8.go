package grammar

import "fmt"

func int8_test() {
	var a int8 = -128
	var b int8 = -1
	c := a / b
	fmt.Printf("c的二进制：%b\n", c)
	println(c) // -128 补码:1000 0000
	// int8[-128~127] -128 / -1 = 128 变量允许溢出，128的补码1000 0000
}
