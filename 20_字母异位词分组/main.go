package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	// 映射关系
	hashMap := make(map[string][]string)
	for _, v := range strs {
		// 转成字符数组
		charArr := []byte(v)
		// 按从小到大排序字符数组
		sort.SliceStable(charArr, func(i, j int) bool {
			return charArr[i] < charArr[j]
		})
		// 将升序排列的字符数组转成字符串，作为 map 的键
		// map 的值为分组之后的字符串，用一个数组保存（ go 切片）
		hashMap[string(charArr)] = append(hashMap[string(charArr)], v)
	}

	// 根据 map 的值生成一个二维数组（go 二维切片）
	for _, v := range hashMap {
		res = append(res, v)
	}
	return res
}


func main() {
	example := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(example))
}
