package main

import "fmt"

func reverse(x int) int {
	var (
		result = 0
	)
	for x != 0 {
		mod := x % 10
		newResult := result*10 + mod
		if (newResult-mod)/10 != result {
			return 0
		}

		result = newResult
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
