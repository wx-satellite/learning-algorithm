package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	// 当前节点为空 或者 当前节点的下一个节点为空 则认为循环结束
	for current != nil && current.Next != nil {
		// 相等，则跳过下一个节点
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			// 不想等，则移动到下一个节点
			current = current.Next
		}
	}
	// 返回
	return head
}

func main() {

	//生成链表
	head := &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next.Next = &ListNode{
		Val: 3,
	}
	head.Next.Next.Next.Next = &ListNode{
		Val: 3,
	}

	head = deleteDuplicates(head)

	// 打印链表
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}



