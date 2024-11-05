package exercise

import (
	"fmt"
	"strconv"
	"testing"
)

func reverse(x int) int {
	flag := false
	if x < 0 {
		x = -x
		flag = true
	}
	// 整形->字符串
	str := strconv.Itoa(x)
	sLen := len(str)
	// 存反转的切片
	sli := make([]byte, sLen)
	for i := 0; i < sLen; i++ {
		sli[sLen-i-1] = str[i]
	}
	x, _ = strconv.Atoi(string(sli))
	if x < 1>>31 || x+1 > 1<<31 {
		return 0
	}
	if flag {
		return -x
	} else {
		return x
	}
}

func Test07(t *testing.T) {
	value := -123
	fmt.Println(reverse(value))
}
