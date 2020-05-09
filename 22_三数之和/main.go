package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) <= 2 {
		return result
	}
	// 需要排序
	sort.Ints(nums)
	a := 0
	// 有三个指针，a 是最左边的，所以 a 最大值为 nums 倒数第三个索引
	for a <= len(nums)- 3 {

		// 保证 a 在向右移动时指向的值，不与之前的值相等
		if a != 0 && nums[a] == nums[a-1]  {
			a++
			continue
		}

		// b 为 a 的后一个指针
		b := a + 1

		// c 指向 nums 最后一个元素
		c := len(nums) -1

		// sum 的值为后续 nums[b] + nums[c] 的值
		sum := 0 - nums[a]

		// b 必须小于 c
		for b < c {

			// 保证 c 在向左偏移时指向的值，不与之前的相等
			if c < len(nums) - 1 && nums[c] == nums[c + 1] {
				c--
				continue
			}

			// 保证 b 在向右偏移时指向的值，不与之前的相等
			if b > a + 1 && nums[b] == nums[b-1] {
				b++
				continue
			}

			// 比较和
			if nums[b]+nums[c] > sum {
				c--

			} else if nums[b]+nums[c] < sum {
				b++

			} else {
				result = append(result, []int{nums[a], nums[b], nums[c]})
				c--
				b++
			}

		}
		// a 向右移动
		a++
	}
	return result

}

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums))
}
