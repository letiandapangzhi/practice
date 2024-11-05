package exercise

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/zigzag-conversion/description/
func convert(s string, numRows int) string {
	sLen := len(s)
	var sIndex, j int

	// 二维数组初始化
	var array [][]byte
	for i := 0; i < numRows; i++ {
		array = append(array, make([]byte, 0))
	}
	isAdd := true
	// 遍历字符串长度
	for sIndex < sLen {
		array[j] = append(array[j], s[sIndex])
		sIndex++

		// 二维数组纵下标是从0～num-1～0，横坐标直接append切片不用管，确保字符在z形转换正确的一层即可
		// 一层永远是0
		if numRows > 1 {
			if isAdd {
				j++
				if j == numRows-1 {
					isAdd = false
				}
			} else {
				j--
				if j == 0 {
					isAdd = true
				}
			}
		}

	}

	data := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		data = append(data, array[i]...)
	}

	return string(data)
}

func Test06(t *testing.T) {
	value := "ab"
	fmt.Println(convert(value, 1))
}
