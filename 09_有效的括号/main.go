package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	stack []byte // 存放字节
	length int // 内部维护的长度
}

// 压栈
func (s *Stack) push(b byte) {
	s.stack = append(s.stack,b)
	s.length ++
}

// 出栈
func (s *Stack) pop() (res byte) {
	if s.length <= 0 {
		return
	}
	s.length --
	res = s.stack[s.length]
	s.stack = s.stack[0:s.length]
	return
}

// 判断栈是否为空
func (s *Stack) isEmpty() bool {
	return s.length <= 0
}

// 构造
func getStack() *Stack {
	return &Stack{
	}
}

func isValid(s string) bool {
	if len(s) <= 0 {
		return true
	}
	// 实例化栈
	stack := getStack()
	for i := 0; i < len(s); i++ {
		// 判断是左括号就压入栈
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack.push(s[i])
		} else {
			// 如果栈为空，这时候 i 没有越界，则返回 false
			if stack.isEmpty() {
				return false
			}
			// 获取栈顶元素
			top := stack.pop()
			// 比较是否匹配
			if '(' == top && s[i] != ')' || '[' == top && s[i] !=']' || '{' == top && s[i] !='}'{
				return false
			}
		}
	}
	// 如果 i 越界，并且 栈 为空，则返回 true
	if stack.isEmpty() {
		return true
	}
	return false
}


func isValidOther(s string) bool {

	// 判断是否有一对相邻的括号
	for strings.Index(s,"()") != -1 || strings.Index(s,"[]") != -1 || strings.Index(s,"{}") != -1 {

		// 存在，则替换成 空字符串， 继续下次判断
		s = strings.Replace(s,"()","",-1)
		s = strings.Replace(s,"[]","",-1)
		s = strings.Replace(s,"{}","",-1)
	}

	// 如果不存在一对相邻的括号，并且剩余的字符串长度不为0，则返回 false
	if len(s) >= 1 {
		return false
	}
	return true
}

func main() {

	fmt.Println(isValidOther("{([])}"))
}
