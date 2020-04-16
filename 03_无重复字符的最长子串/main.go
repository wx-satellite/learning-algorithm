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
		// 字符的ASCII范围在0-127，所以这里申明了map的key类型为uint8
		dict := make(map[uint8]bool)
		dict[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			// 判断字符是否在map中，如果存在就直接跳出循环
			if _, ok := dict[s[j]]; ok {
				break
			}
			// 将字符放入map中
			dict[s[j]] = true
			// 每次都计算最大值
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
	x := 0
	y := 0
	s1 = s[x:y]
	for y < len(s) {
		// strings.IndexByte 用于判断y对应的字符是否在[x,y)中，不存在返回-1
		if index := strings.IndexByte(s1, s[y]); index != -1 {
			// 如果存在的话，x直接跳到index的后一位。
			// 至于为什么是 x += index + 1 而不是 x = index + 1 是因为 index 是根据 s1这个子串获得的， 而 x 的值
			// 在下一步求 s1 的时候是相对于 s原串 的。gi
			x += index + 1
		}
		s1 = s[x : y+1]
		if len(s1) > Length {
			Length = len(s1)
		}
		y++
	}

	return Length
}

func main() {
	var s = "bbtablud"
	fmt.Println(lengthOfLongestSubstringGod(s))
}