package main

import "fmt"


// nums是有序的，并且是从小到大的
func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	i := 0
	j := 1

	for j <= len(nums) -1 {

		// 相等，则 j++
		if nums[j] == nums[i] {
			j++
		} else {
				// 不想等，则 i+1 与 j 指向的元素互换，i 向后移动一位，j 向后移动一位
			i++
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	// 返回长度
	return i + 1
}


func main() {
	s := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(s))
	fmt.Println(s)
}
