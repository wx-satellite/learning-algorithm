package main

import "fmt"

// 二分查找 主要是操作 数组的索引。

// 本题区间 [1, x/2] 因为是连续的正整数，完全可以当作数组的索引使用，所以二分查找可以直接使用，不需要做什么转换

func mySqrt(x int) int {

	left := 1
	right := x / 2


	for left < right {
		// 中间值
		mid := left + (right - left) / 2

		// 如果 中间值的平方 小于x，则移动left到mid的后一位
		if mid * mid < x {
			left = mid + 1
		} else if mid * mid > x {
			// 如果 中间值的平方 大于x，则移动right到mid的前一位
			right = mid - 1
		} else {
			// 相等的话，mid就是x的平方根
			return mid
		}
	}


	// 判断当前left的平方是否大于x
	if left * left > x {
		return left - 1
	}

	return left

}




func main() {
	fmt.Println(mySqrt(9))
}
