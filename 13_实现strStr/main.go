package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	i := 0
	j := 0
	for j < len(haystack) {
		if haystack[j] != needle[i] {
			j = j - i + 1
			i = 0
		} else {
			if i == len(needle) -1 {
				return j - i
			}
			j++
			i++
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("aaaaa","bba"))
}
