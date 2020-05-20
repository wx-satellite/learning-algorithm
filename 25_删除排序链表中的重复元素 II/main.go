package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	origin := &ListNode{Val:0}
	current := head
	pre := origin
	really := origin
	for current != nil {
		// 因为新引入了一个初始节点，防止和原始链表的第一个节点冲突，比如第一个节点刚好 Val 也为 0
		// 那么单纯 current.Val != pre.Val 判断是不够的，需要加入 pre == origin 这个判断
		if (pre == origin || current.Val != pre.Val) && (current.Next == nil || current.Val != current.Next.Val) {
			really.Next = current
			really = current

		}

		pre = current
		current = current.Next

		// 这里需要将 pre 和 current 之间的连接断开，例如链表为： 1---->2---->2，
		// 如果不把 节点1 与 节点2 的连接断开，后续返回的 origin.Next 链表结构为： 1---->2---->2 仔细体会一下！
		pre.Next = nil
	}
	return origin.Next
}


func main() {
	head := &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next = &ListNode{
		Val: 2,
	}
	//head.Next.Next.Next = &ListNode{
	//	Val: 3,
	//}
	//head.Next.Next.Next.Next = &ListNode{
	//	Val: 4,
	//}
	//head.Next.Next.Next.Next.Next = &ListNode{
	//	Val: 4,
	//}
	//head.Next.Next.Next.Next.Next.Next = &ListNode{
	//	Val: 5,
	//}

	res := deleteDuplicates(head)

	for nil != res {
		fmt.Println(res.Val)
		res = res.Next
	}

}
