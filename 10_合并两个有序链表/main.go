package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head = &ListNode{}
		current = head
	)

	for true {

		// l1 为空时，将 current.next 指向 l2 即可
		if l1 == nil {
			current.Next = l2
			break
		}

		// l2 为空时，将 current.next 指向 l1 即可
		if l2 == nil {
			current.Next = l1
			break
		}

		// 如果 l1的值 小于 l2的值 current.next指向l1，l1后移
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			// 如果 l2的值 小于等于 l1的值 current.next指向l2，l2后移
			current.Next = l2
			l2 = l2.Next
		}

		// current 后移一位
		current = current.Next
	}

	// head 指向的是默认的节点，head.next 才是结果链表的头节点
	return head.Next

}


func main() {
	var l1 = &ListNode{
		Val: 1,
	}
	l1.Next = &ListNode{
		Val:2,
	}
	l1.Next.Next = &ListNode{
		Val: 4,
	}


	var l2 = &ListNode{
		Val: 1,
	}
	l2.Next = &ListNode{
		Val:3,
	}
	l2.Next.Next = &ListNode{
		Val: 4,
	}


	res := mergeTwoLists(l1,l2)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
