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
		if height[j] > height[i] {
			minHeight = height[i]
		} else {
			minHeight = height[j]
		}
		newMax := (j - i) * minHeight
		if newMax > max {
			max = newMax
		}

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
