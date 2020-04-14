package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	// 将前缀初始化为第一个元素值
	var prefix = strs[0]
	for i := 1; i< len(strs); i++ {
		// 判断字符串是否是prefix开头的，如果是 strings.Index 返回 0
		if strings.Index(strs[i], prefix) != 0 {
			// 去掉前缀的最后一个字符
			prefix = prefix[:len(prefix)-1]
			// 因为会执行 i++，为了保证 i 不变，先 i--
			i--
		}
	}
	return prefix
}




func main() {
	case1 := []string{"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(case1))
	case2 := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(case2))

}