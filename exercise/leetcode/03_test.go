package exercise

import (
	"strings"
	"testing"
)

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	var num int
	var sub string
	var index int
	for index < len(s) {
		char := s[index : index+1]
		if strings.Contains(sub, char) {
			subIndex := strings.Index(sub, char)
			if len(sub)-subIndex == 1 {
				// 相邻
				sub = char
			} else {
				sub = sub[subIndex+1:] + char
				if num < len(sub) {
					num = len(sub)
				}
			}
		} else {
			sub += char
			if num < len(sub) {
				num = len(sub)
			}
		}
		index++
	}
	return num
}

// 滑动窗口 左右指针
func lengthOfLongestSubstringV1(s string) int {
	num := 0
	sLen := len(s)

	// 映射 map判断窗口子串是否存在重复
	// map[字符]int8
	subMap := make(map[byte]int8)

	// 左右窗口边界
	var right int
	for left := 0; left < sLen; left++ {
		// 左指针滑动，删除子串第一个字符
		if left != 0 {
			delete(subMap, s[left-1])
		}
		// 右指针滑动，直到遍历完毕或子串包含新扩的字符
		// subMap[s[right]] == 0 新包括进来的字符不在子串内
		for right < sLen && subMap[s[right]] == 0 {
			// 标识该字符在子串中
			subMap[s[right]]++
			// 右指针滑动
			right++
		}
		// 包含或者右指针遍历结束,更新最长字串长度
		if num < len(subMap) {
			num = len(subMap)
		}
	}
	return num
}

func Test03(t *testing.T) {
	a := "abcbbae"
	lengthOfLongestSubstringV1(a)
}
