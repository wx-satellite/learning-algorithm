package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	var res string
	var carry uint8
	i := len(a) - 1
	j := len(b) - 1

	// i 或者 j 有一个越界了就跳出循环，即表示 字符串a 与字符串b 长度不一致
	for i >= 0 && j >= 0 {
		// 计算 sum ，strconv.Itoa 将数值转成对应的字符串
		res = strconv.Itoa(int((a[i] - '0' + (b[j] - '0') + carry) % 2)) + res

		// 计算进位
		carry = ((a[i] - '0') + (b[j] - '0') + carry) / 2
		i --
		j --
	}

	// j越界
	for i >= 0 {
		// 字符串a 剩余的字符与进位求和
		res = strconv.Itoa(int((a[i] - '0'  + carry) % 2)) + res
		carry = ((a[i] - '0') + carry) / 2
		i --
	}

	// i 越界
	for j >= 0 {
		// 字符串b 剩余的字符与进位求和
		res = strconv.Itoa(int((b[j] - '0' + carry) % 2)) + res
		carry = ((b[j] - '0') + carry) / 2
		j --
	}

	// i 和 j 都越界了，判断 carry进位是否大于0，如果大于0，这需要将这个最高位补上
	if carry > 0 {
		res = "1" + res
	}
	return  res
}



func main() {
	fmt.Println(addBinary("1010","1011"))
}
