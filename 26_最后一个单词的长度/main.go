package main

import "fmt"

func lengthOfLastWord(s string) int {
	if "" == s {
		return 0
	}
	strLength := len(s)
	result := 0
	for i := strLength - 1; i >= 0; i-- {
		// 遇到空字符串
		if ' ' == s[i]  {
			// 如果 result 大于 0，则表示到了最后一个单词的左边界，返回结果
			if result >0  {
				break
			} else {
				// 如果 result 为 0，则表示目前为止没有遇到过单词，继续从右向左遍历
				continue
			}
		}
		result ++


	}
	return result
}



func main() {
	fmt.Println(lengthOfLastWord("Hello World d "))
}
