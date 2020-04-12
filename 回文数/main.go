package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	div := 1

	// 得到相应位数的最小值
	// 例如：3245，那么div的值就是1000
	for x / (div) >=  10 {
		div *= 10
	}
	for x != 0 {
		// 得到第一位的值
		left := x / div
		// 得到最后一位的值
		right := x % 10
		if left != right {
			return false
		}
		// 去掉第一位和最后一位
		// 例如：（ 3245 - 3 * 1000 ）/ 10 = 24
		x = (x - left * div) / 10

		// 因为去掉了两位，因此div也相应的调整
		div /= 100
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(1000021))
}


