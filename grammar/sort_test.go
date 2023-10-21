package grammar

import (
	"fmt"
	"testing"
)

// 冒泡 两两对比 移动
func sort1(data []int) {
	// 遍历次数
	dataLen := len(data) - 1
	for i := 0; i < dataLen; i++ {
		for j := 0; j < dataLen-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println(data)
}

// 选择 查找最值下标 交换
func sort2(data []int) {
	dataLen := len(data) - 1
	for i := 0; i < dataLen; i++ {
		minNumIndex := i
		for j := i + 1; j <= dataLen; j++ {
			if data[minNumIndex] > data[j] {
				minNumIndex = j
			}
		}
		data[i], data[minNumIndex] = data[minNumIndex], data[i]
	}
	fmt.Println(data)
}

// 插入 从后遍历有序 挪动插入
func sort3(data []int) {
	dataLen := len(data)
	for i := 1; i < dataLen; i++ {
		// 当前待插入值
		compareNum := data[i]
		// 最后一位有序下标
		sortIndex := i - 1
		for sortIndex >= 0 && compareNum < data[sortIndex] {
			// 有序值向后挪动一位
			data[sortIndex+1] = data[sortIndex]
			// 有序下标向前
			sortIndex--
		}
		// 插入值最终确定的位置
		data[sortIndex+1] = compareNum
	}
	fmt.Println(data)
}

func TestSort(t *testing.T) {
	data := []int{3, 1, 2, 4, 5, 6}
	//sort1(data)
	//sort2(data)
	sort3(data)
}
