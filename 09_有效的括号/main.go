package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	stack []byte
	length int
}

func (s *Stack) push(b byte) {
	s.stack = append(s.stack,b)
	s.length ++
}

func (s *Stack) pop() (res byte) {
	if s.length <= 0 {
		return
	}
	s.length --
	res = s.stack[s.length]
	s.stack = s.stack[0:s.length]
	return
}

func (s *Stack) isEmpty() bool {
	return s.length <= 0
}

func getStack() *Stack {
	return &Stack{
	}
}

func isValid(s string) bool {
	if len(s) <= 0 {
		return true
	}
	stack := getStack()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack.push(s[i])
		} else {
			if stack.isEmpty() {
				return false
			}
			top := stack.pop()
			if '(' == top && s[i] != ')' || '[' == top && s[i] !=']' || '{' == top && s[i] !='}'{
				return false
			}
		}
	}
	if stack.isEmpty() {
		return true
	}
	return false
}


func isValidOther(s string) bool {
	for strings.Index(s,"()") != -1 || strings.Index(s,"[]") != -1 || strings.Index(s,"{}") != -1 {
		s = strings.Replace(s,"()","",-1)
		s = strings.Replace(s,"[]","",-1)
		s = strings.Replace(s,"{}","",-1)
	}
	if len(s) >= 1 {
		return false
	}
	return true
}

func main() {

	fmt.Println(isValidOther("{([])}"))
}
