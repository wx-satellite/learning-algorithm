## 删除排序链表中的重复元素 II（Remove-Duplicates-From-Sorted-List-II）
题干如下：
>给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
示例 1:
&nbsp;&nbsp;输入: 1->2->3->3->4->4->5
&nbsp;&nbsp;输出: 1->2->5
示例 2:
&nbsp;&nbsp;输入: 1->1->1->2->3
&nbsp;&nbsp;输出: 2->3

这题是上一篇 [让我们一起啃算法----删除排序链表中的重复元素](https://learnku.com/articles/44475 "让我们一起啃算法----删除排序链表中的重复元素") 的进阶版。在上一篇中，我们删掉了重复的其他元素，仅保留一个元素，例如： **1 -> 2 -> 2 -> 3** 变成 **1 -> 2 -> 3**。这题，则需要将仅保留的一个元素也删掉，例如： **1 -> 2 -> 2 -> 3** 变成 **1 -> 3**。

## 解题思路
借上一篇的思路，我们需要在 current 指针基础上再多准备一个前置指针 pre 和一个有效值指针 really。really 指针始终指向有效链表的尾部，pre 指针指向 current 的前置节点，因为我们需要比较当前节点 current 的值是否和前置节点、后置节点的值相等，在单向链表中，后置节点很容易获取，即 current.next，但是前置节点不能很容易获取到，需要用 pre 记录。

初始化：由于需要一个前置节点，当 current 为链表的头节点 head 时前置节点其实是不存在的，为了方便后续的逻辑，我们需要设置一个初始节点 origin，使 pre、really 指向 origin，current 指向 head。

判断 current 的 val 与 current.next 的 val 和 pre 的 val 是否相等，相等的话，则 pre 指向 current， current 指向 current.next，如果不相等，则 really.next 指向 current，really 指向 current，同时 pre 指向 current，current 指向 current.next，大致思路如上。但是这里有个注意事项：
> 按照上面的逻辑，pre 指向 current 后，需要将 pre.next 置为空，例如链表为： 1 ----> 2 ----> 2，如果不把 节点1 与 节点2 的连接断开，后续返回的 origin.Next 链表结构为： 1 ----> 2 ----> 2 仔细体会一下！

流程图如下：

![](https://cdn.learnku.com/uploads/images/202005/20/21280/GTVFznlUfp.jpg!large)

