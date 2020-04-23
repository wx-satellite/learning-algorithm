package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left

}


func main() {
	fmt.Println(searchInsert([]int{1,3,5,6},4))
}
