package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) -1; i >= 0; i-- {
		// 第一次 加1 操作是题目指定的，后续的 加1 操作是低位的进位
		digits[i] ++

		// 小于10与10相模值不变，等于10与10相模为0，这一步的主要目的就是当存在进位时当前值为 10 需要将其设置为 0
		digits[i] %= 10

		// 如果 不等于0 则表示没有进位，直接返回
		if digits[i] != 0 {
			return digits
		}

		// 等于 0 则表示存在进位，由于上面执行了 digits[i] %= 10 ，这时候当前值已经被设置为 0，后续 i-- 之后，执行 digits[i]++ ，是因为有进位需要加上
	}

	// i 已经越界，并且 digits[0] 为 0 表示存在一个进位 1 ，因此需要对数组扩容
	digits = append([]int{1}, digits...)
	return digits
}


func main() {
	fmt.Println(plusOne([]int{4,3,2,2}))
}
