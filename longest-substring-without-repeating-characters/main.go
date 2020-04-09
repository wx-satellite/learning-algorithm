package main

import (
	"fmt"
	"strings"
)

// 暴力破解
func lengthOfLongestSubstring(s string) (length int) {
	if "" == s {
		return
	}
	length = 1
	for i := 0; i < len(s); i++ {
		set := make(map[uint8]bool)
		set[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			if _, ok := set[s[j]]; ok {
				break
			}
			set[s[j]] = true
			cMax := j - i + 1
			if cMax > length {
				length = cMax
			}
		}
	}
	return
}

/**
滑动窗口：
滑动窗口是数组/字符串问题中常用的抽象概念。
窗口通常是在数组/字符串中由开始和结束索引定义的一系列元素的集合，即 [i, j)（左闭，右开）。
而滑动窗口是可以将两个边界向某一方向“滑动”的窗口。例如，我们将 [i, j) 向右滑动 1 个元素，
则它将变为 [i+1, j+1)[i+1,j+1)（左闭，右开）。
 */
// 优化
func lengthOfLongestSubstringOptimize(s string) (max int) {
	if "" == s {
		return
	}
	var (
		left = -1
		right = 0
		set = make(map[uint8]bool)
	)
	for right < len(s) {
		if _, ok := set[s[right]]; ok {
			left++
			delete(set,s[left])
		} else {
			set[s[right]] = true
			if right - left > max {
				max = right - left
			}
			right ++
		}
	}
	return
}


// 大神解决
func lengthOfLongestSubstringGod(s string) int {
	var Length int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right]
	for ; right < len(s); right++ {

		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > Length {
			Length = len(s1)
		}
	}

	return Length
}

func main() {
	var s = "bbbbb"
	fmt.Println(lengthOfLongestSubstringOptimize(s))
}