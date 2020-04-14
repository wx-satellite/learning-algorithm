package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	var prefix = strs[0]
	for i := 1; i< len(strs); i++ {
		if strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			i--
		}
	}
	return prefix
}




func main() {
	//fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	s := "  hello world!  "
	fmt.Println(reverseWords(s))

}