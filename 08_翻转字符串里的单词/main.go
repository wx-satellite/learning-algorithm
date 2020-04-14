package main

import "strings"

func reverseWords1(s string) string {
	ss := strings.Fields(s)
	var r []string
	for j := len(ss) - 1; j >=0; j-- {
		r = append(r,ss[j])
	}
	return strings.Join(r," ")
}

func reverseWords(s string) string {
	s = strings.Trim(s," ")
	bs := []byte(s)
	var rs []byte
	index := 0
	lastIndex := len(bs)
	for j := len(bs)-1; j >=0;j-- {
		if bs[j] == 32 && index == 0{
			index = j + 1
		}
		if bs[j] == 32 && bs[j-1] != 32 {
			rs = append(rs, bs[index:lastIndex]...)
			rs = append(rs,32)
			lastIndex = j
			index = 0
		}
		if j == 0 {
			rs = append(rs, bs[0:lastIndex]...)
		}
	}
	return string(rs)
}
