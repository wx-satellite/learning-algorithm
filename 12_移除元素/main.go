package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) <= 0 {
		return 0
	}
	i := -1
	for j := 0; j < len(nums); j++ {

		// 如果 j指向的值 和 val 的值不同，那么 i 后移一位，然后 nums[i] 的值 与 nums[j] 的值 互换
		if nums[j] != val {
			i ++
			// 这里其实可以写成：nums[i] = nums[j]
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return i + 1
}




func main() {
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2},2))
}
