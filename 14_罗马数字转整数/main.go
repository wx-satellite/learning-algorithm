package  main

import (
	"fmt"
	"strings"
)


// 罗马数字转成整数
func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i ++ {
		switch s[i] {
		case 'I':
			res += 1
		case 'V':
			res += 5
		case 'X':
			res += 10
		case 'L':
			res += 50
		case 'C':
			res += 100
		case 'D':
			res += 500
		case 'M':
			res += 1000
		}
	}

	if strings.Index(s,"IV") != -1 {
		res -= 2
	}
	if strings.Index(s,"IX") != -1 {
		res -= 2
	}
	if strings.Index(s,"XL") != -1 {
		res -= 20
	}
	if strings.Index(s,"XC") != -1 {
		res -= 20
	}
	if strings.Index(s,"CD") != -1 {
		res -= 200
	}
	if strings.Index(s,"CM") != -1 {
		res -= 200
	}
	return res
}




func getIntByRoman(r byte) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func romanToInt1(s string) int {
	res := 0
	i := 0
	// 跳出循环之后，i刚好指向字符串s的最后一个字符
	for ; i < len(s)-1; i ++ {
		cur := getIntByRoman(s[i])
		next := getIntByRoman(s[i+1])
		// 比较值，如果当前值比下一个值小则减去当前值，反之则加上当前值
		if cur < next {
			res -= cur
		} else {
			res += cur
		}
	}
	// 注意最后一个值需要加上
	res += getIntByRoman(s[i])
	return res
}

func romanToInt2(s string) int {
	res := 0
	i := 0

	// 将上面的函数，改成map
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// 跳出循环之后，i刚好指向字符串s的最后一个字符
	for ; i < len(s)-1; i ++ {
		cur := m[s[i]]
		next := m[s[i+1]]
		// 比较值，如果当前值比下一个值小则减去当前值，反之则加上当前值
		if cur < next {
			res -= cur
		} else {
			res += cur
		}
	}
	// 注意最后一个值需要加上
	res += m[s[i]]
	return res
}



func main() {

	fmt.Println(romanToInt1("D"))
}