package main

import "fmt"

func strStr(haystack string, needle string) int {

	// 子串为空则返回 0
	if needle == "" {
		return 0
	}

	i := 0
	j := 0
	for j < len(haystack) {
		// 判断 haystack[j] 与 needle[i] 是否相等，不想等，则 i 重新指向 needle 字符串的首字符，即 i = 0
		// j 指向 j上一次初始化的后一个字符串，即 j = j - i + 1
		if haystack[j] != needle[i] {
			j = j - i + 1
			i = 0
		} else {
			// haystack[j] 与 needle[i] 相等，则先判断 i 是否已经最大
			// 是的话就返回 needle 在 haystack 的第一个位置，计算方式： j - i
			if i == len(needle) -1 {
				return j - i
			}
			// 如果 i 不是 needle 的最大索引，那么 j 向后移动一位，i 向后移动一位，继续比较
			j++
			i++
		}
	}
	// j 越界了，说明不存在子串 needle，即返回 -1
	return -1
}

func main() {
	fmt.Println(strStr("aaaaa","bba"))
}
