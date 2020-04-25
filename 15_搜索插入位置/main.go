package main

import "fmt"


// 函数的实现并不是最简洁的，但是整个二分法思路是很清晰的。
func searchInsert(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}

	var left, right = 0, len(nums) - 1
	for left <= right {

		// 根据 left 和 right 计算 mid 的值
		mid := left + (right - left) / 2  // ( left + right ) / 2


		// 目标值大于中间值，left = mid + 1
		if nums[mid] < target {
			left = mid + 1

		// 目标值小于中间值 right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// 目标值等于中间值返回 mid，即为目标值需要插入的位置
			return mid
		}
	}

	// 找不到目标值，这时候返回left就是目标值需要插入的位置
	return left

}


func main() {
	fmt.Println(searchInsert([]int{1,3,5,6},4))
}
