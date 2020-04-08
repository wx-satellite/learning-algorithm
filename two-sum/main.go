package main

import "fmt"

/*
	题目：求两数之和
	详情: https://leetcode-cn.com/problems/two-sum/
*/

/*暴力解法
思路：
	两层循环，外部循环的当前值与target的差值为内层循环需要定位的值
*/
func twoSumForce(nums []int, target int) (result []int) {
	length := len(nums)
	if length <= 0 {
		return
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			// 找到了就直接返回
			if nums[j] + nums[i] == target {
				return []int{i,j}
			}
		}
	}
	return
}

/*优雅解法
思路：
	一层循环，引入map，循环的当前值与target的差值在map中查找，如果不存在则将当前值作为键放入map，值的索引作为map的值。
*/
func twoSumElegant(nums []int, target int) (result []int) {
	if len(nums) <= 0 {
		return
	}
	var m = make(map[int]int)
	for index, value := range nums {
		if v, ok := m[target-value]; !ok {
			// 将当前值和索引放入map中
			m[value] = index
		} else {
			return []int{index,v}
		}
	}
	return

}

func main() {
	var (
		nums   = []int{2, 7, 11, 15}
		target = 10
	)

	fmt.Println(twoSumElegant(nums, target))

}
