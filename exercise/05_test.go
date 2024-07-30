package exercise

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/longest-palindromic-substring/description/
func longestPalindrome(s string) string {
	// 最长回文子串
	maxSub := make([]byte, 0)

	sLen := len(s)

	for i := 0; i < sLen; i++ {
		// 剩余部分已经没有机会超过最长子串，提前结束
		if sLen-i <= len(maxSub) {
			break
		}
		// 临时切片子串，判断是否回文
		sub := make([]byte, 0, sLen-i)
		for j := i; j < sLen; j++ {
			sub = append(sub, s[j])
			subLen := len(sub)
			// 遍历子串，判断回文
			// 回文标识
			flag := true
			for subIndex := 0; subIndex < subLen/2; subIndex++ {
				if sub[subIndex] != sub[subLen-1-subIndex] {
					flag = false
					break
				}
			}
			// 更新最长回文子串
			if flag && len(maxSub) < subLen {
				maxSub = sub
			}
		}
	}

	return string(maxSub)
}

// 动态规划（i到j是回文子串，则i+1到j-1一定是回文子串）
// 1. 定义二维数组存i到j是不是回文子串
// 2. 第一层遍历从目标字符串检索符合长度2的回文子串，直到回文子串长度超过目标字符串[L]
// 3. 第二层遍历从目标最左边界[i]开始检索，推算出右边界[j=L-i-1]
// 4. 对比循环的最左最右边界，赋值二维数组标识i到j是否回文子串（i,j = i+1,j-1;由小范围推导至大，小的二维数组一定赋值过了，比如acda不是回文子串）

// 中心扩展（回文中心；奇偶为中心）
func longestPalindromeV1(s string) string {
	sLen := len(s)
	// 定义回文子串开始结束下标
	var start, end int
	// 遍历字符串
	for i := 0; i < sLen; i++ {
		// 分别以奇数/偶数为回文中心,获取该回文子串的左右边界
		left1, right1 := getLeftAndRight(s, sLen, i, i)
		left2, right2 := getLeftAndRight(s, sLen, i, i+1)
		// 获取奇数/偶数回文中心的更长回文子串,更新最长回文子串开始结束下标
		if left1 < right1 && right1-left1 > end-start {
			start = left1
			end = right1
		}
		if left2 < right2 && right2-left2 > end-start {
			start = left2
			end = right2
		}
	}

	return s[start : end+1]
}

func getLeftAndRight(s string, sLen, left, right int) (int, int) {
	for left >= 0 && right < sLen && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func Test05(t *testing.T) {
	value := "aba"
	fmt.Println(longestPalindromeV1(value))
}
