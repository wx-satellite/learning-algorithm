package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		carry        int // 进位
		currentValue int // 要追加节点的Val的，也就是图中的value
		centerValue  int // 中间值
		point        = &ListNode{}
		head         = point
	)
	for l1 != nil && l2 != nil {
		centerValue = l1.Val + l2.Val + carry
		currentValue = centerValue % 10
		carry = centerValue / 10
		// 新增一个节点，用于存储结果
		point.Next = &ListNode{Val: currentValue}
		// point始终指向最后一个节点
		point = point.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	// 上一个for循环执行完毕，有可能l1或者l2仍有节点，下面两个循环就是处理这种情况的

	// 处理两个链表长度不一致的情况
	for l1 != nil {
		centerValue = l1.Val + carry
		currentValue = centerValue % 10
		carry = centerValue / 10
		point.Next = &ListNode{Val: currentValue}
		point = point.Next
		l1 = l1.Next
	}

	// 处理两个链表长度不一致的情况
	for l2 != nil {
		centerValue = l2.Val + carry
		currentValue = centerValue % 10
		carry = centerValue / 10
		point.Next = &ListNode{Val: currentValue}
		point = point.Next
		l2 = l2.Next
	}

	// 如果进位不为0，需再追加一个节点
	if carry == 1 {
		point.Next = &ListNode{Val: 1}
	}

	// 因为head指向的是初始节点，所以head.Next才是我们要的结果链表真正的头节点
	return head.Next
}

func addTwoNumbersOptimize(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		carry        int
		currentValue int
		centerValue  int
		point        = &ListNode{}
		head         = point
	)
	for l1 != nil || l2 != nil || carry != 0 {
		lValue := 0
		rValue := 0
		if l1 != nil {
			lValue = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			rValue = l2.Val
			l2 = l2.Next
		}
		centerValue = lValue + rValue + carry
		currentValue = centerValue % 10
		carry = centerValue / 10
		point.Next = &ListNode{Val: currentValue}
		point = point.Next

	}
	return head.Next
}

func main() {
	leftNode := &ListNode{
		Val: 5,
	}
	leftNode.Next = &ListNode{
		Val: 4,
	}
	leftNode.Next.Next = &ListNode{
		Val: 3,
	}

	rightNode := &ListNode{
		Val: 5,
	}
	rightNode.Next = &ListNode{
		Val: 6,
	}
	rightNode.Next.Next = &ListNode{
		Val: 4,
	}

	result := addTwoNumbersOptimize(leftNode, rightNode)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
