package main

import "fmt"

func reverse(x int) int {
	var (
		result = 0
	)
	// x为正负数不用分类讨论。
	// 因为：-123 % 10 = -3，与正数的处理逻辑一致
	for x != 0 {
		mod := x % 10
		newResult := result*10 + mod
		// 倒推思路： C = A + B，如果没有溢出 A = C - B，溢出之后 A != C - B
		if (newResult-mod)/10 != result {
			return 0
		}

		result = newResult

		// 去掉最后一位
		x = x / 10

	}
	return result
}

func reverseSpecial(x int) int {
	var (
		result int32
		mod int32
	)
	for x != 0 {
		mod = int32(x % 10)
		newResult := result*10 + mod
		if (newResult-mod)/10 != result {
			return 0
		}

		result = newResult
		x = x / 10

	}
	return int(result)
}

func main() {
	fmt.Println(reverseSpecial(1534236469))
}
