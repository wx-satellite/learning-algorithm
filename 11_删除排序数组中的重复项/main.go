package main

import "fmt"


func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	i := 0
	j := 1

	for j <= len(nums) -1 {
		if nums[j] == nums[i] {
			j++
		} else {
			i++
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return i + 1
}


func main() {
	s := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(s))
	fmt.Println(s)
}
