package main

import "fmt"


// 递归
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	return climbStairs(n-1) + climbStairs( n- 2)
}

// 斐波那契数列
// 1 1 2 3 5 8 13 ....
func climbStairs2(n int) int {

	// 1 阶台阶直接返回 1
	if n == 1 {
		return 1
	}

	// 2 阶台阶直接返回 2
	if n == 2 {
		return 2
	}

	current := 2
	pre := 1
	// 当前台阶的走法是前两个台阶走法之和
	for i := 3; i <= n;i ++ {
		current = current + pre
		pre = current - pre
	}
	return current
}


func main() {
	fmt.Println(climbStairs2(44))
}
