package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	// max 指向 nums1 与 nums2 合并之后的最后一个元素
	max := m + n - 1

	// 指向 num1 最后一个元素
	i := m - 1

	// 指向 num2 最后一个元素
	j := n -1


	for i >= 0 && j >= 0 {

		// 从右向左比较值的大小
		if nums1[i] > nums2[j] {
			nums1[max] = nums1[i]

			// i 向左移动
			i--
		} else {
			nums1[max] = nums2[j]

			// j 向左移动
			j--
		}

		// max 向左移动
		max --




	}

	// 如果 i 越界了，将 nums2 剩余的元素赋值到 num1 的 [0,m] 之间
	for j >= 0 {
		nums1[max] = nums2[j]
		max--
		j--
	}

	// 如果 j 越界了，其实下面这个循环可以省略。
	for i >= 0 {
		nums1[max] = nums1[i]
		max--
		i--
	}
}

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	num2 := []int{2, 5, 6}
	n := 3

	merge(num1, m, num2, n)

	fmt.Println(num1)
}
