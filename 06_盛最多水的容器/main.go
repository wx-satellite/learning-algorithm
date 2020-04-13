package main

import "fmt"

func maxArea(height []int) int {
	var (
		i         = 0
		j         = len(height) - 1
		max       = 0
		minHeight = 0
	)

	for i < j {
		// H 的 计算
		if height[j] > height[i] {
			minHeight = height[i]
		} else {
			minHeight = height[j]
		}

		// 记录面积，也就是纳水量
		newMax := (j - i) * minHeight
		if newMax > max {
			max = newMax
		}

		// 移动高度小的指针，如果相等则一起移动
		if height[j] > height[i] {
			i++
		} else if height[j] < height[i] {
			j--
		} else {
			j--
			i++
		}
	}
	return max
}

func main() {

	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(s))
}
